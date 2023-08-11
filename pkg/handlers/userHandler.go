package handlers

import (
	"net/http"
	"github.com/jasperwreed/goliath/pkg/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Handle the request and response here...
}
