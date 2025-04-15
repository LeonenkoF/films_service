package service

import (
	"films_service/internal/entity"
	"films_service/internal/models"
	"films_service/internal/repository"
)

type MovieService struct {
	repoActor repository.Actor
	repoMovie repository.Movie
}

func NewMoiveService(repoActor repository.Actor, repoMovie repository.Movie) *MovieService {
	return &MovieService{
		repoActor: repoActor,
		repoMovie: repoMovie,
	}
}

func (s *MovieService) CreateActor(input models.Actor) (int, error) {
	id, err := s.repoActor.CreateActor(input)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *MovieService) DeleteActor(id int) error {
	if err := s.repoActor.DeleteActor(id); err != nil {
		return err
	}
	return nil
}

func (s *MovieService) UpdateActor(input models.Actor) error {
	err := s.repoActor.UpdateActor(input)
	if err != nil {
		return err
	}
	return nil
}

func (s *MovieService) ActorsList() ([]entity.Movie_actor, error) {
	result, err := s.repoActor.ActorsList()
	if err != nil {
		return nil, err
	}
	return result, nil
}
