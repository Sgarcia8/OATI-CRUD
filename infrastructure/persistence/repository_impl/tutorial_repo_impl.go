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
	_, err := r.ormer.QueryTable(&models_orm.TutorialORM{}).All(&ormList)
	if err != nil {
		return nil, err
	}
	return toTutorialDomainList(ormList), nil
}

func (r *TutorialRepositoryImpl) GetById(id int) (*domain.Tutorial, error) {
	ormTutorial := &models_orm.TutorialORM{Id: id}
	err := r.ormer.Read(ormTutorial)
	if err != nil {
		return nil, err
	}
	_, err = r.ormer.LoadRelated(ormTutorial, "Comments")
	if err != nil {
		return nil, err
	}
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
	_, err := r.ormer.Delete(&models_orm.TutorialORM{Id: id})
	return err
}
