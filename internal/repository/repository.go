package repository

import (
	"spotifyly-ai/internal/models"
	"spotifyly-ai/pkg/spotify"
)

type Repository struct {
	spotifyClient *spotify.Client
}

func NewRepository(spotifyClient *spotify.Client) *Repository {
	return &Repository{spotifyClient: spotifyClient}
}

func (r *Repository) GetLikedSongs() ([]models.Song, error) {
	// Fetch liked songs from Spotify API
	tracks, err := r.spotifyClient.GetLikedSongs()
	if err != nil {
		return nil, err
	}

	// Convert to models.Song
	var songs []models.Song
	for _, track := range tracks {
		songs = append(songs, models.Song{
			Name:   track.Name,
			Artist: track.Artists[0].Name, // Only the first artist is considered
		})
	}

	return songs, nil
}
