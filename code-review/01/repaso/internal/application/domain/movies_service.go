package domain

import (
	"errors"
	"fmt"
)

type Service struct {
}

func (s *Service) FindMovies(title *string, year *uint32, director *string) ([]Movie, error) {
	if year != nil && year < 1900 || year > 2022 {
		return nil, fmt.Errorf("invalid year: %w", ErrInvalidYear)
	}
}

var (
	ErrInvalidYear = errors.New("invalid year")
)

type MovieService interface {
	FindMovies(title *string, year *uint32, director *string) ([]Movie, error)
}