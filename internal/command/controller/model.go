package controller

type PostCreatedDto struct {
	UserName string `json:"user_name"`
	Content  string `json:"content"`
}
