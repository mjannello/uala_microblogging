package service

import (
	"uala/internal/command/eventpublisher"
	"uala/internal/command/eventstore"
	commentmodel "uala/internal/model/comment"
	postmodel "uala/internal/model/post"
	"uala/pkg/clock"
	"uala/pkg/logger"
)

type CommandService interface {
	AddPost(userName, content string) (postmodel.PostAddedEvent, error)
	UpdatePost(userName, content string, id int64) (postmodel.PostUpdatedEvent, error)
	DeletePost(userName string, id int64) (postmodel.PostDeletedEvent, error)
	AddCommentToPost(postID int64, userName, content string) (commentmodel.CommentAddedEvent, error)
	//UpdateComment(model post.PostUpdatedEvent) (comment.CommentUpdatedEvent, error)
	//DeleteComment(model post.PostDeletedEvent) (comment.CommentDeletedEvent, error)
}

type commandService struct {
	clock                clock.Clock
	eventStoreRepository eventstore.EventStore
	eventPublisher       eventpublisher.EventPublisher
}

func NewCommandService(store eventstore.EventStore, publisher eventpublisher.EventPublisher, clock clock.Clock) CommandService {
	return &commandService{
		clock:                clock,
		eventStoreRepository: store,
		eventPublisher:       publisher,
	}
}

func (s *commandService) AddPost(userName, content string) (postmodel.PostAddedEvent, error) {
	logger.Logger.Printf("add post")

	createdPostEvent := postmodel.NewPostAddedEvent(userName, content, s.clock.Time())

	logger.Logger.Printf("save event")
	eventID, err := s.eventStoreRepository.SaveEvent(createdPostEvent)

	if err != nil {
		logger.Logger.Printf("error storing model", err)
		return postmodel.PostAddedEvent{}, err
	}

	createdPostEvent.ID = eventID
	if err := s.eventPublisher.Publish(createdPostEvent); err != nil {
		logger.Logger.Printf("error publishing model", err)
		return postmodel.PostAddedEvent{}, err
	}

	return createdPostEvent, nil
}

func (s *commandService) UpdatePost(userName, content string, id int64) (postmodel.PostUpdatedEvent, error) {
	logger.Logger.Printf("update post")

	updatedPostEvent := postmodel.NewPostUpdatedEvent(userName, content, id, s.clock.Time())

	eventID, err := s.eventStoreRepository.SaveEvent(updatedPostEvent)

	if err != nil {
		logger.Logger.Printf("error storing model", err)
		return postmodel.PostUpdatedEvent{}, err
	}

	updatedPostEvent.ID = eventID
	if err := s.eventPublisher.Publish(updatedPostEvent); err != nil {
		logger.Logger.Printf("error publishing model", err)
		return postmodel.PostUpdatedEvent{}, err
	}

	return updatedPostEvent, nil
}

func (s *commandService) DeletePost(userName string, id int64) (postmodel.PostDeletedEvent, error) {
	logger.Logger.Printf("delete post")

	deletedPostEvent := postmodel.NewPostDeletedEvent(userName, id, s.clock.Time())

	eventID, err := s.eventStoreRepository.SaveEvent(deletedPostEvent)

	if err != nil {
		logger.Logger.Printf("error storing model", err)
		return postmodel.PostDeletedEvent{}, err
	}

	deletedPostEvent.ID = eventID
	if err := s.eventPublisher.Publish(deletedPostEvent); err != nil {
		logger.Logger.Printf("error publishing model", err)
		return postmodel.PostDeletedEvent{}, err
	}

	return deletedPostEvent, nil
}

func (s *commandService) AddCommentToPost(postID int64, userName, content string) (commentmodel.CommentAddedEvent, error) {
	logger.Logger.Printf("add comment to postID")

	createdCommentEvent := commentmodel.NewCommentAddedEvent(postID, userName, content, s.clock.Time())

	logger.Logger.Printf("save event")
	eventID, err := s.eventStoreRepository.SaveEvent(createdCommentEvent)

	if err != nil {
		logger.Logger.Printf("error storing model", err)
		return commentmodel.CommentAddedEvent{}, err
	}

	createdCommentEvent.ID = eventID
	if err := s.eventPublisher.Publish(createdCommentEvent); err != nil {
		logger.Logger.Printf("error publishing model", err)
		return commentmodel.CommentAddedEvent{}, err
	}

	return createdCommentEvent, nil
}
