package controllers

import (
	"net/http"
	"bugounty/internals/services"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	userService *services.UserService
}


func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func UserController(handler *UserHandler) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/create", handler.CreateHandler).Methods(http.MethodPost)
	return router
}


func (h *UserHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	
}
