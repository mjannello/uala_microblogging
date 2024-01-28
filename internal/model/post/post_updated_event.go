package post

import "time"

type PostUpdatedEvent struct {
	ID            int64
	UserName      string
	Content       string
	PostUpdatedID int64
	DateCreated   time.Time
	Type          string
}

func NewPostUpdatedEvent(userName, content string, postUpdatedID int64, timeNow time.Time) PostUpdatedEvent {
	return PostUpdatedEvent{
		Type:          "PostUpdatedEvent",
		UserName:      userName,
		Content:       content,
		PostUpdatedID: postUpdatedID,
		DateCreated:   timeNow,
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

func (pue PostUpdatedEvent) GetData() map[string]interface{} {
	return map[string]interface{}{
		"UserName":      pue.UserName,
		"PostDeletedID": pue.PostUpdatedID,
		"Content":       pue.Content,
		"DateCreated":   pue.DateCreated,
	}
}
