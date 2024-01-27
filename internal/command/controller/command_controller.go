package controller

import (
	"encoding/json"
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
	var requestData PostCreatedDto
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	createdPost, err := cc.commandService.AddPost(requestData.UserName, requestData.Content)
	if err != nil {
		http.Error(w, "Error processing AddPost command", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdPost)
	if err != nil {
		return
	}
}

func (cc *commandController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement UpdatePost
	return
}

func (cc *commandController) DeletePost(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement DeletePost
	return
}
