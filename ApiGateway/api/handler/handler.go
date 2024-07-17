package handlers

import (
	"item_api/genproto"
)

type Handler struct {
	AuthService genproto.AuthServiceClient
	EcoService  genproto.EcoServiceClient
}

func NewHandler(authService genproto.AuthServiceClient, ecoService genproto.EcoServiceClient) *Handler {
	return &Handler{
		AuthService: authService,
		EcoService:  ecoService,
	}
}
