package service

import (
	"uala/internal/command/eventpublisher"
	"uala/internal/command/eventstore"
	post2 "uala/internal/model/post"
	"uala/pkg/clock"
	"uala/pkg/logger"
)

type CommandService interface {
	AddPost(userID, content string) (post2.PostAddedEvent, error)
	UpdatePost(event post2.PostUpdatedEvent) (post2.PostUpdatedEvent, error)
	DeletePost(event post2.PostDeletedEvent) (post2.PostDeletedEvent, error)
	// TODO:
	//AddComment(model post.PostAddedEvent) (comment.CommentAddedEvent, error)
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

func (s *commandService) AddPost(userID, content string) (post2.PostAddedEvent, error) {
	logger.Logger.Printf("add post")

	createdPostEvent := post2.NewPostAddedEvent(userID, content, s.clock.Time())

	logger.Logger.Printf("save event")
	eventID, err := s.eventStoreRepository.SaveEvent(createdPostEvent)

	if err != nil {
		logger.Logger.Printf("error storing model", err)
		return post2.PostAddedEvent{}, err
	}

	createdPostEvent.ID = eventID
	logger.Logger.Print("just to publish event")
	if err := s.eventPublisher.Publish(createdPostEvent); err != nil {
		logger.Logger.Printf("error publishing model", err)
		return post2.PostAddedEvent{}, err
	}

	return createdPostEvent, nil
}

func (s *commandService) UpdatePost(event post2.PostUpdatedEvent) (post2.PostUpdatedEvent, error) {
	// TODO: Implement UpdatePost
	return post2.PostUpdatedEvent{}, nil
}

func (s *commandService) DeletePost(event post2.PostDeletedEvent) (post2.PostDeletedEvent, error) {
	// TODO: Implement DeletePost
	return post2.PostDeletedEvent{}, nil
}
