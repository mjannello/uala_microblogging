package post

import "time"

type PostUpdatedEvent struct {
	ID          int64
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

func (pue PostUpdatedEvent) GetType() string {
	return pue.Type
}

func (pue PostUpdatedEvent) GetContent() string {
	return pue.Content
}

func (pue PostUpdatedEvent) GetUserName() string {
	return pue.UserName
}

func (pue PostUpdatedEvent) GetDate() time.Time {
	return pue.DateCreated
}
