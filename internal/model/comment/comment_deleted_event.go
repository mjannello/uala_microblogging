package comment

import "time"

type CommentDeletedEvent struct {
	ID               int64
	UserName         string
	Content          string
	CommentDeletedID int64
	DateCreated      time.Time
	Type             string
}

func NewCommentDeletedEvent(userName string, commentDeletedID int64, timeNow time.Time) CommentDeletedEvent {
	return CommentDeletedEvent{
		Type:             "CommentDeletedEvent",
		UserName:         userName,
		CommentDeletedID: commentDeletedID,
		DateCreated:      timeNow,
	}
}

func (pde CommentDeletedEvent) GetType() string {
	return pde.Type
}

func (pde CommentDeletedEvent) GetContent() string {
	return pde.Content
}

func (pde CommentDeletedEvent) GetUserName() string {
	return pde.UserName
}

func (pde CommentDeletedEvent) GetDate() time.Time {
	return pde.DateCreated
}

func (pde CommentDeletedEvent) GetData() map[string]interface{} {
	return map[string]interface{}{
		"UserName":         pde.UserName,
		"CommentDeletedID": pde.CommentDeletedID,
		"DateCreated":      pde.DateCreated,
	}
}
