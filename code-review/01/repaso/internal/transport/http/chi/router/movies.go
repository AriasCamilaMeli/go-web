package router

import (
	"regexp"

	"github.com/alkemy/repaso/internal/application/domain"
	"github.com/alkemy/repaso/internal/infrastructure/persistance/inmemory"
	"github.com/alkemy/repaso/internal/transport/http/chi/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func NewChiMoviesRouter() *chi.Mux {
	db := map[uuid.UUID]domain.Movie{
		uuid.New(): {
			ID:          uuid.New(),
			Title:       "The Godfather",
			Director:    "Francis Ford Coppola",
			ReleaseYear: 1972,
		},
		uuid.New(): {
			ID:          uuid.New(),
			Title:       "The Shawshank Redemption",
			Director:    "Frank Darabont",
			ReleaseYear: 1972,
		},
		uuid.New(): {
			ID:          uuid.New(),
			Title:       "The Dark Knight",
			Director:    "Frank Darabont",
			ReleaseYear: 1998,
		},
	}
	rp := inmemory.NewMoviesMap(nil)
	hd := handlers.NewChiMovies(rp)

	rt := chi.NewRouter()
	rt.Get("/", hd.Get)

	return rt
}


// email is a value object
// - it is immutable
// - it is validated
type email struct {
	value string
}

func (e email) GetValue() string {
	return e.value
}

func NewEmail(value string) (*email, error) {
	// regex
	if !regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, value) {
		return nil, domain.ErrInvalidEmail
	}
	return &email{value: value}, nil
}
