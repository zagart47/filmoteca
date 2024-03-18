package entity

type Movie struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"release_date"`
	Rating      float32 `json:"rating"`
}

func NewMovie() Movie {
	return Movie{}
}

type Options struct {
	Field  string
	Order  string
	Search string
	Actor  string
	Title  string
}
