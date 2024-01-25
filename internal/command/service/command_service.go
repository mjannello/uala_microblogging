package service

import (
	"uala/internal/command/event/post"
	"uala/internal/command/eventpublisher"
	"uala/internal/command/repository"
)

type CommandService interface {
	AddPost(event post.PostAddedEvent) (post.PostAddedEvent, error)
	UpdatePost(event post.PostUpdatedEvent) (post.PostUpdatedEvent, error)
	DeletePost(event post.PostDeletedEvent) (post.PostDeletedEvent, error)
	// TODO:
	//AddComment(event post.PostAddedEvent) (comment.CommentAddedEvent, error)
	//UpdateComment(event post.PostUpdatedEvent) (comment.CommentUpdatedEvent, error)
	//DeleteComment(event post.PostDeletedEvent) (comment.CommentDeletedEvent, error)
}

type commandService struct {
	EventStoreRepository repository.EventStore
	EventPublisher       eventpublisher.EventPublisher
}

func NewCommandService(store repository.EventStore, publisher eventpublisher.EventPublisher) CommandService {
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
