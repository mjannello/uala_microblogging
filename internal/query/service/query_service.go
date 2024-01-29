package service

import (
	"errors"
	"fmt"
	"time"
	"uala/internal/query/model/feed"
	"uala/internal/query/repository"
	"uala/pkg/logger"
)

type QueryService interface {
	GetFeed() (feed.Feed, error)
	UpdateRepositoryWithEvent(map[string]interface{}) error
}

type queryService struct {
	queryRepository repository.QueryRepository
}

func NewQueryService(r repository.QueryRepository) QueryService {
	return &queryService{
		queryRepository: r,
	}
}

func (qs *queryService) GetFeed() (feed.Feed, error) {
	return qs.queryRepository.GetFeed()

}

func (qs *queryService) UpdateRepositoryWithEvent(eventData map[string]interface{}) error {
	logger.Logger.Print("eventData", eventData)

	eventType, ok := eventData["Type"].(string)
	if !ok {
		return errors.New("type is not a string")
	}
	var err error
	switch eventType {
	case "PostAddedEvent":
		err = qs.handlePostAddedEvent(eventData)
	case "PostDeletedEvent":
		err = qs.handlePostDeletedEvent(eventData)
	case "PostUpdatedEvent":
		err = qs.handlePostUpdatedEvent(eventData)
	case "CommentAddedEvent":
		err = qs.handleCommentAddedEvent(eventData)
	default:
		err = fmt.Errorf("unknown event type: %s", eventType)
	}
	if err != nil {
		return err
	}
	return nil
}

func (qs *queryService) handlePostAddedEvent(eventData map[string]interface{}) error {
	logger.Logger.Print(eventData)
	postIDFloat, ok := eventData["ID"].(float64)
	if !ok {
		return fmt.Errorf("invalid post ID")
	}

	postIDInt := int64(postIDFloat)

	postCreated := feed.Post{
		ID:          postIDInt,
		UserName:    eventData["UserName"].(string),
		Content:     eventData["Content"].(string),
		DateCreated: parseDateString(eventData["DateCreated"].(string)),
		Comments:    []feed.Comment{},
	}
	_, err := qs.queryRepository.SavePost(postCreated)
	if err != nil {
		return err
	}
	return nil
}

func (qs *queryService) handlePostDeletedEvent(eventData map[string]interface{}) error {
	logger.Logger.Print(eventData)
	postToDeleteIDFloat, ok := eventData["PostDeletedID"].(float64)
	if !ok {
		return fmt.Errorf("invalid post to delete ID")
	}
	postToDeleteIDInt := int64(postToDeleteIDFloat)

	userName := eventData["UserName"].(string)
	postDeletedID := postToDeleteIDInt
	_, err := qs.queryRepository.DeletePost(userName, postDeletedID)
	if err != nil {
		return err
	}
	return nil
}

func (qs *queryService) handlePostUpdatedEvent(eventData map[string]interface{}) error {
	logger.Logger.Print(eventData)
	postToUpdateIDFloat, ok := eventData["PostUpdatedID"].(float64)
	if !ok {
		return fmt.Errorf("invalid post to update ID")
	}
	postToUpdateIDInt := int64(postToUpdateIDFloat)

	postUpdated := feed.Post{
		ID:      postToUpdateIDInt,
		Content: eventData["Content"].(string),
	}

	userName := eventData["UserName"].(string)
	err := qs.queryRepository.UpdatePost(userName, postUpdated)
	if err != nil {
		return err
	}
	return nil
}

func (qs *queryService) handleCommentAddedEvent(eventData map[string]interface{}) error {
	logger.Logger.Print(eventData)
	commentIDFloat, ok := eventData["ID"].(float64)
	if !ok {
		return fmt.Errorf("invalid comment ID")
	}
	commentIDInt := int64(commentIDFloat)

	postIDFloat, ok := eventData["PostID"].(float64)
	if !ok {
		return fmt.Errorf("invalid post ID")
	}
	postIDInt := int64(postIDFloat)

	comment := feed.Comment{
		ID:          commentIDInt,
		PostID:      postIDInt,
		UserName:    eventData["UserName"].(string),
		Content:     eventData["Content"].(string),
		DateCreated: parseDateString(eventData["DateCreated"].(string)),
	}

	return qs.queryRepository.AddCommentToPost(postIDInt, comment)
}

func parseDateString(dateString string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339Nano, dateString)
	if err != nil {
		logger.Logger.Printf("Error parsing date: %v", err)
	}
	return parsedTime
}
