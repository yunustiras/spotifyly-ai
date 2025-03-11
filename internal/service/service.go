package service

import (
	"spotifyly-ai/internal/repository"
	"spotifyly-ai/pkg/ai"
)

type Song struct {
	Name   string
	Artist string
}

type Service struct {
	repo     *repository.Repository
	aiClient *ai.AIClient
}

func NewService(repo *repository.Repository, aiClient *ai.AIClient) *Service {
	return &Service{repo: repo, aiClient: aiClient}
}

func (s *Service) GroupSongsByCriteria(criteria string) (map[string][]Song, error) {
	songs, err := s.repo.GetLikedSongs()
	if err != nil {
		return nil, err
	}

	// Group songs by criteria using AI
	groupedSongs, err := s.aiClient.GroupSongs(songs, criteria)
	if err != nil {
		return nil, err
	}

	return groupedSongs, nil
}
