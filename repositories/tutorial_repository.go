package repositories

import "oati-crud-comentarios/domain"

type TutorialRepository interface {
	GetAll() ([]*domain.Tutorial, error)
	GetById(id int) (*domain.Tutorial, error)
	Create(tutorial *domain.Tutorial) error
	Update(tutorial *domain.Tutorial) error
	Delete(id int) error
}
