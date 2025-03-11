package main

import (
	"fmt"
	"log"
	"spotifyly-ai/internal/handler"
	"spotifyly-ai/internal/repository"
	"spotifyly-ai/internal/service"
	"spotifyly-ai/pkg/ai"
	"spotifyly-ai/pkg/spotify"
)

func main() {
	// Initialize Spotify client
	spotifyClient, err := spotify.NewClient()
	if err != nil {
		log.Fatalf("Failed to initialize Spotify client: %v", err)
	}

	// Initialize AI client
	aiClient, err := ai.NewAIClient()
	if err != nil {
		log.Fatalf("Failed to initialize AI client: %v", err)
	}

	// Initialize repository, service, and handler
	repo := repository.NewRepository(spotifyClient)
	svc := service.NewService(repo, aiClient)
	h := handler.NewHandler(svc)

	// Fetch and group songs
	groupedSongs, err := h.GroupSongsByCriteria("genre") // Example criteria: "genre"
	if err != nil {
		log.Fatalf("Failed to group songs: %v", err)
	}

	// Print grouped songs
	for group, songs := range groupedSongs {
		fmt.Printf("Group: %s\n", group)
		for _, song := range songs {
			fmt.Printf(" - %s by %s\n", song.Name, song.Artist)
		}
	}
}
