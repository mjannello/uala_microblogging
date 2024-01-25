package controller

import (
	"net/http"
	"uala/internal/command/service"
)

type CommandController interface {
	AddPost(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

type commandController struct {
	commandService service.CommandService
}

func NewCommandController(cs service.CommandService) CommandController {
	return &commandController{
		commandService: cs,
	}
}

func (cc *commandController) AddPost(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement AddPost
	return
}

func (cc *commandController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement UpdatePost
	return
}

func (cc *commandController) DeletePost(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement DeletePost
	return
}
