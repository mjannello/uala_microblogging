package comment

import "time"

type CommentAddedEvent struct {
	ID          int64
	PostID      int64
	UserName    string
	Content     string
	DateCreated time.Time
	Type        string
}

func NewCommentAddedEvent(postID int64, userName, content string, timeNow time.Time) CommentAddedEvent {
	return CommentAddedEvent{
		Type:        "CommentAddedEvent",
		PostID:      postID,
		UserName:    userName,
		Content:     content,
		DateCreated: timeNow,
	}
}

func (cae CommentAddedEvent) GetType() string {
	return cae.Type
}

func (cae CommentAddedEvent) GetContent() string {
	return cae.Content
}

func (cae CommentAddedEvent) GetUserName() string {
	return cae.UserName
}

func (cae CommentAddedEvent) GetDate() time.Time {
	return cae.DateCreated
}

func (cae CommentAddedEvent) GetPostID() int64 {
	return cae.PostID
}

func (cae CommentAddedEvent) GetData() map[string]interface{} {
	return map[string]interface{}{
		"PostID":      cae.PostID,
		"UserName":    cae.UserName,
		"Content":     cae.Content,
		"DateCreated": cae.DateCreated,
	}
}
