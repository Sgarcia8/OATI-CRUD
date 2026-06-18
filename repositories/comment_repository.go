package repositories

import "oati-crud-comentarios/domain"

type CommentRepository interface {
	GetByTutorialId(tutorialId int) ([]*domain.Comment, error)
	GetById(id int) (*domain.Comment, error)
	Create(comment *domain.Comment) error
	Update(comment *domain.Comment) error
	Delete(id int) error
}
