package model

type Movie struct {
	popularity       float32 `json:"popularity"`
	voteCount        int     `json:"vote_count"`
	video            bool    `json:"video"`
	posterPath       string  `json:"poster_path"`
	id               int     `json:"id"`
	adult            bool    `json:"adult"`
	backdropPath     string  `json:"backdrop_path"`
	originalLanguage string  `json:"original_language"`
	originalTitle    string  `json:"original_title"`
	genreIds         []int   `json:"genre_ids"`
	title            string  `json:"title"`
	voteAverage      float32 `json:"vote_average"`
	overview         string  `json:"overview"`
	releaseDate      string  `json:"release_date"`
}
