package post

import "time"

type PostDeletedEvent struct {
	ID          int64
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
	// TODO: implement
	return nil
}
