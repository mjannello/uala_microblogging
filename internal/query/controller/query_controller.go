package controller

import (
	"encoding/json"
	"net/http"
	"uala/internal/query/service"
)

type QueryController interface {
	GetFeed(w http.ResponseWriter, r *http.Request)
	GetFeedByUser(w http.ResponseWriter, r *http.Request)
}

type queryController struct {
	queryService service.QueryService
}

func NewQueryController(s service.QueryService) QueryController {
	return &queryController{
		queryService: s,
	}
}

func (qc *queryController) GetFeed(w http.ResponseWriter, r *http.Request) {
	feed, err := qc.queryService.GetFeed()
	if err != nil {
		http.Error(w, "Error getting feed", http.StatusInternalServerError)
		return
	}
	feedDTO := FeedModelToRest(feed)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feedDTO)
}

func (qc *queryController) GetFeedByUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user_name")

	userFeed, err := qc.queryService.GetFeedByUser(username)
	if err != nil {
		http.Error(w, "Error getting user feed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userFeed)
}
