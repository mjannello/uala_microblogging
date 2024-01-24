package service

import (
	"uala/internal/command/event/post"
	"uala/internal/command/eventpublisher"
	"uala/internal/command/repository"
	"uala/pkg/event"
)

type CommandService interface {
	AddPost(event post.PostAddedEvent) (post.PostAddedEvent, error)
	UpdatePost(event post.PostUpdatedEvent) (post.PostUpdatedEvent, error)
	DeletePost(event post.PostDeletedEvent) (post.PostDeletedEvent, error)
}

type commandService struct {
	EventStoreRepository repository.EventStore
	EventPublisher       eventpublisher.EventPublisher
}

func NewCommandService(store repository.EventStore, publisher event.EventPublisher) CommandService {
	return &commandService{
		EventStoreRepository: store,
		EventPublisher:       publisher,
	}
}

func (s *commandService) AddPost(event post.PostAddedEvent) (post.PostAddedEvent, error) {
	// TODO: Implement AddPost
	return post.PostAddedEvent{}, nil
}

func (s *commandService) UpdatePost(event post.PostUpdatedEvent) (post.PostUpdatedEvent, error) {
	// TODO: Implement UpdatePost
	return post.PostUpdatedEvent{}, nil
}

func (s *commandService) DeletePost(event post.PostDeletedEvent) (post.PostDeletedEvent, error) {
	// TODO: Implement DeletePost
	return post.PostDeletedEvent{}, nil
}
