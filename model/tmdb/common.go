package tmdb

type ApiConfig struct {
	Key string `envconfig:"API_KEY"`
}

type Base struct {
	Page         int64          `json:"page"`
	TotalResults int64          `json:"total_results"`
	TotalPages   int64          `json:"total_pages"`
	Rusults      []BaseContents `json:"results"`
}

type BaseContents struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	PosterPath  string `json:"poster_path"`
}
