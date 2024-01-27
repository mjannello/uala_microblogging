package post

import "time"

type PostDeletedEvent struct {
	ID          string
	UserName    string
	Content     string
	DateCreated time.Time
	Type        string
}

func NewPostDeletedEvent(userID, content string) PostDeletedEvent {
	return PostDeletedEvent{
		UserName: userID,
		Content:  content,
	}
}

func (pde PostDeletedEvent) EventType() string {
	return pde.Type
}
