package controller

type PostCreatedDto struct {
	Content string `json:"content"`
}

type PostUpdatedDto struct {
	Content string `json:"new_content"`
}
