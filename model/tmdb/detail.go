package tmdb

type MovieDetail struct {
	ID               int64   `json:"id"`
	Title            string  `json:"title"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	HomePage         string  `json:"home_page"`
	Genres           []Genre `json:"genres"`
	Overview         string  `json:"overview"`
	VoteAverage      int64   `json:"vote_average"`
	VoteCount        int64   `json:"vote_count"`
}

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
