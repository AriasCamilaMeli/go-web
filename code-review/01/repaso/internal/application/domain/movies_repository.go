package domain

import "github.com/google/uuid"

// type Optional[T any] struct {
// 	value T
// 	exist bool
// }

type Query struct {
	// id Optional[uuid.UUID]
	Title *string

	Director *string
	
	ReleaseYear *releaseYear
}

func NewQuery(title, director *string, releaseYear *releaseYear) *Query {
	return &Query{
		Title: title,
		Director: director,
		ReleaseYear: releaseYear,
	}
}

type MovieRepository interface {
	// Search retrieves all movies, although if queries are present, it filters those registries
	Search(q Query) (map[uuid.UUID]Movie, error)
}