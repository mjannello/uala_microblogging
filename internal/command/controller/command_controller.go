package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"uala/internal/command/service"
)

type CommandController interface {
	AddPost(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
	AddCommentToPost(w http.ResponseWriter, r *http.Request)
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
	userName := r.Header.Get("user_name")
	var requestData PostCreatedDto
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	createdPost, err := cc.commandService.AddPost(userName, requestData.Content)
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
	userName := r.Header.Get("user_name")

	vars := mux.Vars(r)

	postIDStr := vars["id"]

	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid post ID", http.StatusBadRequest)
		return
	}
	var requestData PostUpdatedDto
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	updatedPost, err := cc.commandService.UpdatePost(userName, requestData.Content, postID)
	if err != nil {
		http.Error(w, "Error processing UpdatePost command", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(updatedPost)
	if err != nil {
		return
	}
}

func (cc *commandController) DeletePost(w http.ResponseWriter, r *http.Request) {
	userName := r.Header.Get("user_name")

	vars := mux.Vars(r)

	postIDStr := vars["id"]

	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid post ID", http.StatusBadRequest)
		return
	}

	deletedPost, err := cc.commandService.DeletePost(userName, postID)
	if err != nil {
		http.Error(w, "Error processing DeletePost command", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	err = json.NewEncoder(w).Encode(deletedPost)
	if err != nil {
		return
	}
}

func (cc *commandController) AddCommentToPost(w http.ResponseWriter, r *http.Request) {
	postIDStr := mux.Vars(r)["post_id"]
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid post ID", http.StatusBadRequest)
		return
	}
	userName := r.Header.Get("user_name")

	var requestData CommentAddedDto
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	createdComment, err := cc.commandService.AddCommentToPost(postID, userName, requestData.Content)
	if err != nil {
		http.Error(w, "Error processing AddCommentToPost command", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdComment)
	if err != nil {
		// Handle encoding error
		return
	}
}
