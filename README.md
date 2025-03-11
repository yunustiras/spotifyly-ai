# Spotifyly AI

Spotifyly AI is a Go application that fetches your liked songs from Spotify and groups them based on specific criteria using AI. This application uses the OpenAI GPT API to analyze songs and create customized playlists.

## Features
- Fetches liked songs from Spotify.
- Groups songs based on specific criteria (e.g., genre, artist, mood, year, etc.).
- Uses AI to analyze songs and generate customized playlists.

## Installation

### Clone the repository:
```bash
git clone https://github.com/yunustiras/spotifyly-ai.git
```

### Navigate to the project directory:
```bash
cd spotifyly-ai
```

### Install the required dependencies:
```bash
go mod download
```

### Build the project:
```bash
go build -o spotifyly-ai cmd/spotifyly-ai/main.go
```

### Run the application:
```bash
./spotifyly-ai
```

## Usage
When you run the application, it fetches your liked songs from Spotify and groups them based on the specified criteria. You can pass the grouping criteria as a parameter to the `GroupSongsByCriteria` function.

### Example Criteria

#### Genre:
```go
groupedSongs, err := h.GroupSongsByCriteria("genre")
```
##### Example Output:
```go
map[string][]Song{
    "Pop": {
        {Name: "Shape of You", Artist: "Ed Sheeran"},
        {Name: "Blinding Lights", Artist: "The Weeknd"},
    },
    "Rock": {
        {Name: "Bohemian Rhapsody", Artist: "Queen"},
    },
}
```

#### Artist:
```go
groupedSongs, err := h.GroupSongsByCriteria("artist")
```
##### Example Output:
```go
map[string][]Song{
    "Ed Sheeran": {
        {Name: "Shape of You", Artist: "Ed Sheeran"},
    },
    "Queen": {
        {Name: "Bohemian Rhapsody", Artist: "Queen"},
    },
}
```

#### Mood:
```go
groupedSongs, err := h.GroupSongsByCriteria("mood")
```
##### Example Output:
```go
map[string][]Song{
    "Energetic": {
        {Name: "Uptown Funk", Artist: "Mark Ronson ft. Bruno Mars"},
    },
    "Sad": {
        {Name: "Someone Like You", Artist: "Adele"},
    },
}
```

#### Year:
```go
groupedSongs, err := h.GroupSongsByCriteria("year")
```
##### Example Output:
```go
map[string][]Song{
    "2010s": {
        {Name: "Rolling in the Deep", Artist: "Adele"},
    },
    "2020s": {
        {Name: "Levitating", Artist: "Dua Lipa"},
    },
}
```

## How It Works

### Fetching Songs
The application fetches your liked songs from Spotify using the Spotify API.

### Grouping Songs
The fetched songs are sent to the OpenAI GPT API, which analyzes and groups them based on the specified criteria.

### Displaying Results
The grouped songs are displayed in the terminal.

## License
This project is licensed under the MIT License.
