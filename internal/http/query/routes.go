package http

import (
	"github.com/gorilla/mux"
	"uala/internal/query/controller"
)

type RouterHandler struct {
	QueryController controller.QueryController
}

func NewRouterHandler(qc controller.QueryController) *RouterHandler {
	return &RouterHandler{
		QueryController: qc,
	}
}

func (rh *RouterHandler) ConfigureRoutes(router *mux.Router) {
	rh.routeURLs(router)
}

func (rh *RouterHandler) routeURLs(router *mux.Router) {
	router.HandleFunc("/api/feed", rh.QueryController.GetFeed).Methods("GET")
	router.HandleFunc("/api/feed/{user_name}", rh.QueryController.GetFeedByUser).Methods("GET")

}
