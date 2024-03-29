package repository

import "uala/internal/query/model/feed"

type QueryRepository interface {
	GetFeed() (feed.Feed, error)
	SavePost(post feed.Post) (string, error)
	UpdatePost(userName string, postUpdated feed.Post) error
	DeletePost(userName string, postDeletedID int64) (int64, error)
	AddCommentToPost(postID int64, comment feed.Comment) error
}
