package post

import "time"

type PostDeletedEvent struct {
	ID            int64
	UserName      string
	Content       string
	PostDeletedID int64
	DateCreated   time.Time
	Type          string
}

func NewPostDeletedEvent(userName string, postDeletedID int64, timeNow time.Time) PostDeletedEvent {
	return PostDeletedEvent{
		Type:          "PostDeletedEvent",
		UserName:      userName,
		PostDeletedID: postDeletedID,
		DateCreated:   timeNow,
	}
}

func (pde PostDeletedEvent) GetType() string {
	return pde.Type
}

func (pde PostDeletedEvent) GetContent() string {
	return pde.Content
}

func (pde PostDeletedEvent) GetUserName() string {
	return pde.UserName
}

func (pde PostDeletedEvent) GetDate() time.Time {
	return pde.DateCreated
}

func (pde PostDeletedEvent) GetData() map[string]interface{} {
	return map[string]interface{}{
		"UserName":      pde.UserName,
		"PostDeletedID": pde.PostDeletedID,
		"DateCreated":   pde.DateCreated,
	}
}
