package services

import (
	"errors"
	"oati-crud-comentarios/domain"
	"oati-crud-comentarios/repositories"
)

var (
	ErrCommentNotFound            = errors.New("comentario no encontrado")
	ErrCommentInvalidData         = errors.New("el contenido del comentario es obligatorio")
	ErrTutorialNotFoundForComment = errors.New("el tutorial asociado no existe")
)

type CommentService struct {
	commentRepo  repositories.CommentRepository
	tutorialRepo repositories.TutorialRepository
}

func NewCommentService(
	commentRepo repositories.CommentRepository,
	tutorialRepo repositories.TutorialRepository,
) *CommentService {
	return &CommentService{
		commentRepo:  commentRepo,
		tutorialRepo: tutorialRepo,
	}
}

func (s *CommentService) GetByTutorialId(tutorialId int) ([]*domain.Comment, error) {
	_, err := s.tutorialRepo.GetById(tutorialId)
	if err != nil {
		return nil, ErrTutorialNotFoundForComment
	}
	return s.commentRepo.GetByTutorialId(tutorialId)
}

func (s *CommentService) Create(content string, tutorialId int) (*domain.Comment, error) {
	_, err := s.tutorialRepo.GetById(tutorialId)
	if err != nil {
		return nil, ErrTutorialNotFoundForComment
	}
	comment := domain.NewComment(content, tutorialId)
	if !comment.IsValid() {
		return nil, ErrCommentInvalidData
	}
	if err := s.commentRepo.Create(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *CommentService) Update(id int, content string) (*domain.Comment, error) {
	comment, err := s.commentRepo.GetById(id)
	if err != nil {
		return nil, ErrCommentNotFound
	}
	comment.Update(content)
	if !comment.IsValid() {
		return nil, ErrCommentInvalidData
	}
	if err := s.commentRepo.Update(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *CommentService) Delete(id int) error {
	_, err := s.commentRepo.GetById(id)
	if err != nil {
		return ErrCommentNotFound
	}
	return s.commentRepo.Delete(id)
}
