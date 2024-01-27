package post

import "time"

type PostUpdatedEvent struct {
	ID          string
	UserName    string
	Content     string
	DateCreated time.Time
	Type        string
}

func NewPostUpdatedEvent(userID, content string) PostUpdatedEvent {
	return PostUpdatedEvent{
		UserName: userID,
		Content:  content,
	}
}

func (pue PostUpdatedEvent) EventType() string {
	return pue.Type
}
