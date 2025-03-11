package spotify

import (
	"context"
	"fmt"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

type Client struct {
	client *spotify.Client
}

func NewClient() (*Client, error) {
	// Spotify API credentials
	clientID := "your-client-id"
	clientSecret := "your-client-secret"
	redirectURL := "http://localhost:8080/callback"

	// OAuth2 configuration
	auth := spotifyauth.New(
		spotifyauth.WithClientID(clientID),
		spotifyauth.WithClientSecret(clientSecret),
		spotifyauth.WithRedirectURL(redirectURL),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserLibraryRead,
		),
	)

	// Get the OAuth2 token
	token, err := auth.Exchange(context.Background(), "authorization-code")
	if err != nil {
		return nil, fmt.Errorf("failed to get OAuth2 token: %v", err)
	}

	// Create Spotify client
	httpClient := auth.Client(context.Background(), token)
	client := spotify.New(httpClient)

	return &Client{client: client}, nil
}

func (c *Client) GetLikedSongs() ([]spotify.FullTrack, error) {
	// Fetch liked songs from Spotify API
	results, err := c.client.CurrentUsersTracks(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch liked songs: %v", err)
	}

	// Convert to []spotify.FullTrack
	var tracks []spotify.FullTrack
	for _, item := range results.Tracks {
		tracks = append(tracks, item.FullTrack)
	}

	return tracks, nil
}
