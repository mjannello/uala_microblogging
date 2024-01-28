package repository

import "uala/internal/query/model/feed"

type QueryRepository interface {
	GetFeed() (feed.Feed, error)
	GetFeedByUser(userName string) (feed.Feed, error)
	SavePost(post feed.Post) (string, error)
}
