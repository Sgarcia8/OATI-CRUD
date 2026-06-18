package repository_impl

import (
	"oati-crud-comentarios/domain"
	"oati-crud-comentarios/infrastructure/persistence/models_orm"

	"github.com/beego/beego/v2/client/orm"
)

type TutorialRepositoryImpl struct {
	ormer orm.Ormer
}

func NewTutorialRepository() *TutorialRepositoryImpl {
	return &TutorialRepositoryImpl{ormer: orm.NewOrm()}
}

func (r *TutorialRepositoryImpl) GetAll() ([]*domain.Tutorial, error) {
	var ormList []*models_orm.TutorialORM
	_, err := r.ormer.QueryTable(&models_orm.TutorialORM{}).
		Filter("IsDeleted", false).
		All(&ormList)
	if err != nil {
		return nil, err
	}
	return toTutorialDomainList(ormList), nil
}

func (r *TutorialRepositoryImpl) GetById(id int) (*domain.Tutorial, error) {
	ormTutorial := &models_orm.TutorialORM{}
	err := r.ormer.QueryTable(&models_orm.TutorialORM{}).
		Filter("Id", id).
		Filter("IsDeleted", false).
		One(ormTutorial)
	if err != nil {
		return nil, err
	}

	var ormComments []*models_orm.CommentORM
	_, err = r.ormer.QueryTable(&models_orm.CommentORM{}).
		Filter("Tutorial__Id", id).
		Filter("IsDeleted", false).
		All(&ormComments)
	if err != nil {
		return nil, err
	}
	ormTutorial.Comments = ormComments

	return toTutorialDomain(ormTutorial), nil
}

func (r *TutorialRepositoryImpl) Create(tutorial *domain.Tutorial) error {
	ormTutorial := toTutorialORM(tutorial)
	id, err := r.ormer.Insert(ormTutorial)
	if err != nil {
		return err
	}
	tutorial.Id = int(id)
	return nil
}

func (r *TutorialRepositoryImpl) Update(tutorial *domain.Tutorial) error {
	ormTutorial := toTutorialORM(tutorial)
	_, err := r.ormer.Update(ormTutorial, "Title", "Description", "PublishedAt", "UpdatedAt")
	return err
}

func (r *TutorialRepositoryImpl) Delete(id int) error {
	tutorial, err := r.GetById(id)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		return err
	}

	_, err = tx.QueryTable(&models_orm.CommentORM{}).
		Filter("Tutorial__Id", id).
		Filter("IsDeleted", false).
		Update(orm.Params{"IsDeleted": true})
	if err != nil {
		tx.Rollback()
		return err
	}

	tutorial.SoftDelete()
	ormTutorial := toTutorialORM(tutorial)
	_, err = tx.Update(ormTutorial, "IsDeleted", "UpdatedAt")
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
