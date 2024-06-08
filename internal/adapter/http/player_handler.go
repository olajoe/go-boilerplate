package http

// Contains the implementation details for external dependencies (HTTP handlers)

import (
	"go-boilerplate/internal/ports"
	"net/http"
)

type PlayerHandler struct {
	service ports.PlayerService
}

func NewPlayerHandler(service ports.PlayerService) *PlayerHandler {
	return &PlayerHandler{service: service}
}

func (h *PlayerHandler) GetPlayerByID(w http.ResponseWriter, r *http.Request) {
	// Handle HTTP GET request
}
