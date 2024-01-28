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
	GetFeedByUser(userName string) (feed.Feed, error)
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

func (qs *queryService) GetFeedByUser(userName string) (feed.Feed, error) {
	return qs.queryRepository.GetFeedByUser(userName)
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

	postCreated := feed.Post{
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

func (qs *queryService) handleCommentAddedEvent(data map[string]interface{}) error {
	return nil
}
func parseDateString(dateString string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339Nano, dateString)
	if err != nil {
		logger.Logger.Printf("Error parsing date: %v", err)
	}
	return parsedTime
}