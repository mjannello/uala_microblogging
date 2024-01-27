package post

import "time"

type PostAddedEvent struct {
	ID          string
	UserName    string
	Content     string
	DateCreated time.Time
	Type        string
}

func NewPostAddedEvent(userName, content string, timeNow time.Time) PostAddedEvent {
	return PostAddedEvent{
		Type:        "PostAddedEvent",
		UserName:    userName,
		Content:     content,
		DateCreated: timeNow,
	}
}

func (pae PostAddedEvent) GetType() string {
	return pae.Type
}

func (pae PostAddedEvent) GetContent() string {
	return pae.Content
}

func (pae PostAddedEvent) GetUserName() string {
	return pae.UserName
}

func (pae PostAddedEvent) GetDate() time.Time {
	return pae.DateCreated
}
