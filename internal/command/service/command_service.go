package service

import (
	"uala/internal/command/event/post"
	"uala/internal/command/eventpublisher"
	"uala/internal/command/repository"
	"uala/pkg/clock"
	"uala/pkg/logger"
)

type CommandService interface {
	AddPost(userID, content string) (post.PostAddedEvent, error)
	UpdatePost(event post.PostUpdatedEvent) (post.PostUpdatedEvent, error)
	DeletePost(event post.PostDeletedEvent) (post.PostDeletedEvent, error)
	// TODO:
	//AddComment(event post.PostAddedEvent) (comment.CommentAddedEvent, error)
	//UpdateComment(event post.PostUpdatedEvent) (comment.CommentUpdatedEvent, error)
	//DeleteComment(event post.PostDeletedEvent) (comment.CommentDeletedEvent, error)
}

type commandService struct {
	Clock                clock.Clock
	EventStoreRepository repository.EventStore
	EventPublisher       eventpublisher.EventPublisher
}

func NewCommandService(store repository.EventStore, publisher eventpublisher.EventPublisher, clock clock.Clock) CommandService {
	return &commandService{
		Clock:                clock,
		EventStoreRepository: store,
		EventPublisher:       publisher,
	}
}

func (s *commandService) AddPost(userID, content string) (post.PostAddedEvent, error) {
	logger.Logger.Printf("add post")

	createdPostEvent := post.NewPostAddedEvent(userID, content, s.Clock.Time())

	if err := s.EventPublisher.Publish(createdPostEvent); err != nil {
		logger.Logger.Printf("error publishing event", err)
		return post.PostAddedEvent{}, err
	}
	if err := s.EventStoreRepository.SaveEvent(createdPostEvent); err != nil {
		logger.Logger.Printf("error storing event", err)
		return post.PostAddedEvent{}, err
	}

	return createdPostEvent, nil
}

func (s *commandService) UpdatePost(event post.PostUpdatedEvent) (post.PostUpdatedEvent, error) {
	// TODO: Implement UpdatePost
	return post.PostUpdatedEvent{}, nil
}

func (s *commandService) DeletePost(event post.PostDeletedEvent) (post.PostDeletedEvent, error) {
	// TODO: Implement DeletePost
	return post.PostDeletedEvent{}, nil
}
