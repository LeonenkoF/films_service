package service

import (
	"films_service/internal/entity"
	"films_service/internal/models"
)

func (s *MovieService) CreateMovie(input models.Movie) (int, error) {
	id, err := s.repoMovie.CreateMovie(input)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *MovieService) UpdateMovie(input models.Movie) error {
	err := s.repoMovie.UpdateMovie(input)
	if err != nil {
		return err
	}
	return nil
}

func (s *MovieService) DeleteMovie(id int) error {
	if err := s.repoMovie.DeleteMovie(id); err != nil {
		return err
	}
	return nil
}

func (s *MovieService) MoviesList(sortedBy string, direct string) ([]entity.Movie, error) {
	result, err := s.repoMovie.MoviesList(sortedBy, direct)
	if err != nil {
		return nil, err
	}
	return result, nil
}
