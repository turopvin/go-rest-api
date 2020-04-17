package model

//is used for client response
type ResponseMovie struct {
	Title       string
	ReleaseDate string
}

type TmdbMovie struct {
	Popularity       float32 `json:"popularity"`
	VoteCount        int     `json:"vote_count"`
	Video            bool    `json:"video"`
	PosterPath       string  `json:"poster_path"`
	Id               int     `json:"id"`
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	GenreIds         []int   `json:"genre_ids"`
	Title            string  `json:"title"`
	VoteAverage      float32 `json:"vote_average"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
}

type ChannelMovie struct {
	ApiName string          `json:"api_name"`
	Movies  []ResponseMovie `json:"movies"`
}
