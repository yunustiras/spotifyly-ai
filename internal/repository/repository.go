package repository

import (
	"spotifyly-ai/internal/service"
	"spotifyly-ai/pkg/spotify"
)

type Repository struct {
	spotifyClient *spotify.Client
}

func NewRepository(spotifyClient *spotify.Client) *Repository {
	return &Repository{spotifyClient: spotifyClient}
}

func (r *Repository) GetLikedSongs() ([]service.Song, error) {
	// Fetch liked songs from Spotify API
	songs, err := r.spotifyClient.GetLikedSongs()
	if err != nil {
		return nil, err
	}

	// Convert to service.Song
	var result []service.Song
	for _, s := range songs {
		result = append(result, service.Song{
			Name:   s.Name,
			Artist: s.Artists[0].Name, // Only the first artist is considered
		})
	}

	return result, nil
}
