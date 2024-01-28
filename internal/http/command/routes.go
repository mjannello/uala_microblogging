package http

import (
	"github.com/gorilla/mux"
	commandcontroller "uala/internal/command/controller"
)

type RouterHandler struct {
	CommandController commandcontroller.CommandController
}

func NewRouterHandler(cc commandcontroller.CommandController) *RouterHandler {
	return &RouterHandler{
		CommandController: cc,
	}
}

func (rh *RouterHandler) ConfigureRoutes(router *mux.Router) {
	rh.routeURLs(router)
}

func (rh *RouterHandler) routeURLs(router *mux.Router) {
	router.HandleFunc("/api/post", rh.CommandController.AddPost).Methods("POST")
	router.HandleFunc("/api/post/{id}", rh.CommandController.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/post/{id}", rh.CommandController.DeletePost).Methods("DELETE")
}
