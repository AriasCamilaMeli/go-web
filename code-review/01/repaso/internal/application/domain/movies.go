package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	//
)

type genre uint8

const (
	Horror genre = iota
	Comedy
	SciFi
	Romantic
	Action
)

type releaseYear struct {
	value uint32
}

func (r *releaseYear) Get() uint32 {
	return r.value
}

func NewReleaseYear(year uint32) (*releaseYear, error) {
	if year > uint32(time.Now().Year()) {
		return &releaseYear{}, errors.New("invalid year, must be previous or current to this year")
	}

	return &releaseYear{
		value: year,
	}, nil
}

type Movie struct {
	// id is the unique identifier of a movie
	id uuid.UUID

	// title is the name of the movie
	title string

	// releaseYear is the year where the movie was released
	releaseYear releaseYear

	// director
	director string

	// genre
	genre genre
}

func (m Movie) GetTitle() string {
	return m.title
}

func (m Movie) GetReleaseYear() releaseYear {
	return m.releaseYear
}

func (m Movie) GetDirector() string {
	return m.director
}

func NewMovie(id uuid.UUID, title string, releaseYear releaseYear, director string, genre genre) *Movie {
	if id == uuid.Nil {
		id = uuid.New()
	}

	return &Movie{
		id: id,
		title: title,
		releaseYear: releaseYear,
		director: director,
		genre: genre,
	}
}