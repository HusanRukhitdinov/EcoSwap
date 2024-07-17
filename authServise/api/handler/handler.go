package handlers

import (
	"eco_system/service"
)

type Handler struct {
	// AuthService genproto.AuthServiceClient
	AuthService *service.UserService
}

func NewHandler(auth *service.UserService) *Handler {
	return &Handler{
		AuthService: auth,
		// AuthService: authService,
	}
}
