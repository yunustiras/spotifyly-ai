package handler

import (
	"spotifyly-ai/internal/models"
	"spotifyly-ai/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) GroupSongsByCriteria(criteria string) (map[string][]models.Song, error) {
	return h.service.GroupSongsByCriteria(criteria)
}
