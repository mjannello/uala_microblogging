package feed

import "time"

type Feed struct {
	Posts []Post
}

type Post struct {
	ID          int64
	UserName    string
	Content     string
	DateCreated time.Time
	Comments    []Comment
}

type Comment struct {
	ID          string
	UserName    string
	Content     string
	DateCreated time.Time
	Reactions   []Reaction
}

type Reaction struct {
	ID    string
	Emoji string
}
