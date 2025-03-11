package spotify

import (
	"context"
	"fmt"

	"github.com/zmb3/spotify/v2"
)

type Client struct {
	client *spotify.Client
}

func NewClient() (*Client, error) {
	// Spotify api credentials
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
	token, err := auth.Token(context.Background(), "user-state")
	if err != nil {
		return nil, fmt.Errorf("failed to get OAuth2 token: %v", err)
	}

	// Create spotify client
	httpClient := auth.Client(context.Background(), token)
	client := spotify.New(httpClient)

	return &Client{client: client}, nil
}

func (c *Client) GetLikedSongs() ([]spotify.FullTrack, error) {
	// Fetch liked songs from spotify api
	results, err := c.client.CurrentUsersTracks(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch liked songs: %v", err)
	}

	return results.Tracks, nil
}
