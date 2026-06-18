package repository_impl

import (
	"oati-crud-comentarios/domain"
	"oati-crud-comentarios/infrastructure/persistence/models_orm"

	"github.com/beego/beego/v2/client/orm"
)

type CommentRepositoryImpl struct {
	ormer orm.Ormer
}

func NewCommentRepository() *CommentRepositoryImpl {
	return &CommentRepositoryImpl{ormer: orm.NewOrm()}
}

func (r *CommentRepositoryImpl) GetByTutorialId(tutorialId int) ([]*domain.Comment, error) {
	var ormList []*models_orm.CommentORM
	_, err := r.ormer.QueryTable(&models_orm.CommentORM{}).
		Filter("Tutorial__Id", tutorialId).
		RelatedSel().
		All(&ormList)
	if err != nil {
		return nil, err
	}
	return toCommentDomainList(ormList), nil
}

func (r *CommentRepositoryImpl) GetById(id int) (*domain.Comment, error) {
	ormComment := &models_orm.CommentORM{Id: id}
	err := r.ormer.Read(ormComment)
	if err != nil {
		return nil, err
	}
	return toCommentDomain(ormComment), nil
}

func (r *CommentRepositoryImpl) Create(comment *domain.Comment) error {
	ormComment := toCommentORM(comment)
	id, err := r.ormer.Insert(ormComment)
	if err != nil {
		return err
	}
	comment.Id = int(id)
	return nil
}

func (r *CommentRepositoryImpl) Update(comment *domain.Comment) error {
	ormComment := toCommentORM(comment)
	_, err := r.ormer.Update(ormComment, "Content", "UpdatedAt")
	return err
}

func (r *CommentRepositoryImpl) Delete(id int) error {
	_, err := r.ormer.Delete(&models_orm.CommentORM{Id: id})
	return err
}
