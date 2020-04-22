package model

//is used for client response
type ResponseMovie struct {
	Title       string
	ReleaseDate string
	Director    string
	ImdbRating  string
	Actors      string
	Description string
	TrailerLink string
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
	TrailerLink      string  `json:"trailer_link"`
}

type TmdbMovieVideos struct {
	Id         string `json:"id"`
	Iso639_1   string `json:"iso_639_1"`
	Iso3166_1  string `json:"iso_3166_1"`
	YoutubeKey string `json:"key"`
	Name       string `json:"name"`
	Site       string `json:"site"`
	Size       string `json:"size"`
	Type       string `json:"type"`
}

type OmdbMovie struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	ImdbRating string `json:"imdbRating"`
}

type ChannelMovie struct {
	ApiName string          `json:"api_name"`
	Movies  []ResponseMovie `json:"movies"`
}
