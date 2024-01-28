package controller

import (
	"time"
	feed "uala/internal/query/model/feed"
)

type FeedDTO struct {
	Posts []PostDTO `json:"posts"`
}

type PostDTO struct {
	UserName    string       `json:"user_name"`
	Content     string       `json:"content"`
	DateCreated time.Time    `json:"date_created"`
	Comments    []CommentDTO `json:"comments"`
}

type CommentDTO struct {
	UserName    string        `json:"user_name"`
	Content     string        `json:"content"`
	DateCreated time.Time     `json:"date_created"`
	Reactions   []ReactionDTO `json:"reactions"`
}

type ReactionDTO struct {
	Emoji string `json:"emoji"`
}

func FeedModelToRest(f feed.Feed) FeedDTO {
	var postsDTO []PostDTO
	for _, post := range f.Posts {
		postDTO := PostModelToRest(post)
		postsDTO = append(postsDTO, postDTO)
	}

	return FeedDTO{Posts: postsDTO}
}

func PostModelToRest(post feed.Post) PostDTO {
	var commentsDTO []CommentDTO
	for _, comment := range post.Comments {
		commentDTO := CommentModelToRest(comment)
		commentsDTO = append(commentsDTO, commentDTO)
	}

	return PostDTO{
		UserName:    post.UserName,
		Content:     post.Content,
		DateCreated: post.DateCreated,
		Comments:    commentsDTO,
	}
}

func CommentModelToRest(comment feed.Comment) CommentDTO {
	var reactionsDTO []ReactionDTO
	for _, reaction := range comment.Reactions {
		reactionDTO := ReactionModelToRest(reaction)
		reactionsDTO = append(reactionsDTO, reactionDTO)
	}
	return CommentDTO{
		UserName:    comment.UserName,
		Content:     comment.Content,
		DateCreated: comment.DateCreated,
		Reactions:   reactionsDTO,
	}
}

func ReactionModelToRest(reaction feed.Reaction) ReactionDTO {
	return ReactionDTO{
		Emoji: reaction.Emoji,
	}
}
