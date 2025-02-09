package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/alkemy/repaso/internal/application/domain"
)

func NewChiMovies(rp domain.MovieRepository) *ChiMovies {
	return &ChiMovies{
		rp: rp,
	}
}

type ChiMovies struct {
	rp domain.MovieRepository
}

func (c *ChiMovies) Get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	
	var title *string
	if queryParams.Has("title") {
		titleParam := queryParams.Get("title")
		title = &titleParam
	}

	var director *string
	if queryParams.Has("director") {
		directorStr := queryParams.Get("director")
		director = &directorStr
	}

	releaseYear, _ := domain.NewReleaseYear(0)
	if queryParams.Has("release_year") {
		yearParam, err := strconv.ParseUint(queryParams.Get("release_year"), 2, 32)
		if err != nil {
			http.Error(w, "invalid year, must be numeric and positive", http.StatusBadRequest)
			return
		}
		releaseYear, err = domain.NewReleaseYear(uint32(yearParam))
		if err != nil {
			http.Error(w, "invalid year", http.StatusUnprocessableEntity)
			return
		}
	}

	query := domain.NewQuery(title, director, releaseYear)

	// call repo
	movies, err := c.rp.Search(*query)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidYear):
			http.Error(w, "invalid year", http.StatusUnprocessableEntity)
		default:
			http.Error(w, "query failed", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"data": movies,
	})
}