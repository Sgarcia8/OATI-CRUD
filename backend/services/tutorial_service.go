package services

import (
	"errors"
	"oati-crud-comentarios/domain"
	"oati-crud-comentarios/repositories"
	"time"
)

var (
	ErrTutorialNotFound    = errors.New("tutorial no encontrado")
	ErrTutorialInvalidData = errors.New("título y descripción son obligatorios")
)

type TutorialService struct {
	repo repositories.TutorialRepository
}

func NewTutorialService(repo repositories.TutorialRepository) *TutorialService {
	return &TutorialService{repo: repo}
}

func (s *TutorialService) GetAll() ([]*domain.Tutorial, error) {
	return s.repo.GetAll()
}

func (s *TutorialService) GetById(id int) (*domain.Tutorial, error) {
	tutorial, err := s.repo.GetById(id)
	if err != nil {
		return nil, ErrTutorialNotFound
	}
	return tutorial, nil
}

func (s *TutorialService) Create(title, description string, publishedAt time.Time) (*domain.Tutorial, error) {
	tutorial := domain.NewTutorial(title, description, publishedAt)
	if !tutorial.IsValid() {
		return nil, ErrTutorialInvalidData
	}
	if err := s.repo.Create(tutorial); err != nil {
		return nil, err
	}
	return tutorial, nil
}

func (s *TutorialService) Update(id int, title, description string, publishedAt time.Time) (*domain.Tutorial, error) {
	tutorial, err := s.repo.GetById(id)
	if err != nil {
		return nil, ErrTutorialNotFound
	}
	tutorial.Update(title, description, publishedAt)
	if !tutorial.IsValid() {
		return nil, ErrTutorialInvalidData
	}
	if err := s.repo.Update(tutorial); err != nil {
		return nil, err
	}
	return tutorial, nil
}

func (s *TutorialService) Delete(id int) error {
	_, err := s.repo.GetById(id)
	if err != nil {
		return ErrTutorialNotFound
	}
	return s.repo.Delete(id)
}
