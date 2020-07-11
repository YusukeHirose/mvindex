package model

type ApiConfig struct {
	Key string `envconfig:"API_KEY"`
}

type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	PosterPath  string `json:"poster_path"`
}
