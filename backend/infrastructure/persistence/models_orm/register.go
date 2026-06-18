package models_orm

import "github.com/beego/beego/v2/client/orm"

func RegisterModels() {
	orm.RegisterModel(
		&TutorialORM{},
		&CommentORM{},
	)
}
