package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"uala/internal/query/model/feed"
	"uala/internal/query/repository"
	"uala/pkg/logger"
)

const (
	dbName          = "uala_query_db"
	postsCollection = "posts"
)

type mongoDBRepository struct {
	client *mongo.Client
}

func NewMongoDBRepository(connectionString string) (repository.QueryRepository, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("could not connect to MongoDB: %w", err)
	}

	return &mongoDBRepository{client: client}, nil
}

func (mr *mongoDBRepository) GetFeed() (feed.Feed, error) {
	logger.Logger.Print("start getting feed")
	collection := mr.client.Database(dbName).Collection(postsCollection)
	logger.Logger.Print("collection", collection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logger.Logger.Print("cancel", cancel)

	filter := bson.D{}

	cursor, err := collection.Find(ctx, filter)
	logger.Logger.Print("cursor", cursor)

	if err != nil {
		return feed.Feed{}, fmt.Errorf("error al obtener las publicaciones: %w", err)
	}
	defer cursor.Close(ctx)

	var posts []feed.Post
	for cursor.Next(ctx) {
		logger.Logger.Print("next")

		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return feed.Feed{}, fmt.Errorf("error al decodificar el documento: %w", err)
		}
		logger.Logger.Print("result", result)

		post := feed.Post{
			ID:          result["id"].(int64),
			UserName:    result["username"].(string),
			Content:     result["content"].(string),
			DateCreated: result["datecreated"].(primitive.DateTime).Time(),
		}

		comments := mr.extractComments(result)

		post.Comments = comments

		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		return feed.Feed{}, fmt.Errorf("error during traversing docs: %w", err)
	}

	feed := feed.Feed{
		Posts: posts,
	}

	return feed, nil
}

func (mr *mongoDBRepository) extractComments(result bson.M) []feed.Comment {
	var comments []feed.Comment
	commentsArrayRaw := result["comments"].(primitive.A)
	for _, commentMapRaw := range commentsArrayRaw {
		m := commentMapRaw.(primitive.M)
		comment := feed.Comment{
			ID:      m["id"].(int64),
			Content: m["content"].(string),
		}
		comments = append(comments, comment)

	}
	return comments
}

func (mr *mongoDBRepository) GetFeedByUser(userName string) (feed.Feed, error) {
	return feed.Feed{}, nil
}

func (mr *mongoDBRepository) SavePost(post feed.Post) (string, error) {
	collection := mr.client.Database(dbName).Collection(postsCollection)

	result, err := collection.InsertOne(context.Background(), post)
	logger.Logger.Print("result", result)

	if err != nil {
		return "", fmt.Errorf("could not save post to MongoDB: %w", err)
	}
	insertedID, ok := result.InsertedID.(string)
	if !ok {
		return "", fmt.Errorf("unexpected type for inserted ID")
	}

	return insertedID, nil
}

func (mr *mongoDBRepository) UpdatePost(userName string, postUpdated feed.Post) error {
	collection := mr.client.Database(dbName).Collection(postsCollection)

	filter := bson.D{
		{"username", userName},
		{"id", postUpdated.ID},
	}

	update := bson.D{
		{"$set", bson.D{
			{"content", postUpdated.Content},
		}},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return fmt.Errorf("could not update post in MongoDB: %w", err)
	}

	return nil
}

func (mr *mongoDBRepository) DeletePost(userName string, postDeletedID int64) (int64, error) {
	collection := mr.client.Database(dbName).Collection(postsCollection)

	filter := bson.D{
		{"username", userName},
		{"id", postDeletedID},
	}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, fmt.Errorf("error al eliminar el post: %w", err)
	}

	if result.DeletedCount == 0 {
		return 0, fmt.Errorf("there was no post ID %d for the user %s", postDeletedID, userName)
	}

	return postDeletedID, nil
}

func (mr *mongoDBRepository) AddCommentToPost(postID int64, comment feed.Comment) error {
	collection := mr.client.Database(dbName).Collection(postsCollection)

	filter := bson.M{"id": postID}

	update := bson.M{
		"$push": bson.M{"comments": comment},
	}

	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Logger.Print("fallo al actualizar")
		return fmt.Errorf("error adding comment to post: %w", err)
	}
	logger.Logger.Print(updated)

	return nil
}
