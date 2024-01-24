package controller

import (
	"net/http"
	"uala/internal/command/service"
)

type CommandController interface {
	AddPost(w http.ResponseWriter, r *http.Request) error
	UpdatePost(w http.ResponseWriter, r *http.Request) error
	DeletePost(w http.ResponseWriter, r *http.Request) error
}

type commandController struct {
	PostCommandService *service.CommandService
}

func (cc *commandController) AddPost(w http.ResponseWriter, r *http.Request) error {
	// TODO: Implement AddPost
	return nil
}

func (cc *commandController) UpdatePost(w http.ResponseWriter, r *http.Request) error {
	// TODO: Implement UpdatePost
	return nil
}

func (cc *commandController) DeletePost(w http.ResponseWriter, r *http.Request) error {
	// TODO: Implement DeletePost
	return nil
}
