package inmemory

import (
	"github.com/alkemy/repaso/internal/application/domain"
	"github.com/google/uuid"
)

type MoviesMap struct {
	db map[uuid.UUID]domain.Movie
}

// Search implements domain.MovieRepository.
func (m *MoviesMap) Search(q domain.Query) (map[uuid.UUID]domain.Movie, error) {
	filtered := make(map[uuid.UUID]domain.Movie)

	for key, value := range m.db {
		if q.Title != nil && *q.Title != value.GetTitle() {
			continue
		}
		if q.ReleaseYear != nil && *q.ReleaseYear != value.GetReleaseYear() {
			continue
		}
		if q.Director != nil && *q.Director != value.GetDirector() {
			continue
		}

		// default behaviour
		filtered[key] = value
	}

	return filtered, nil
}

func NewMoviesMap(db map[uuid.UUID]domain.Movie) domain.MovieRepository {
	return &MoviesMap{
		db: db,
	}
}
