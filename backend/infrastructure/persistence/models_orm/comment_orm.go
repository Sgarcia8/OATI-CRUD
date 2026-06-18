package models_orm

import (
	"time"
)

type CommentORM struct {
	Id        int          `orm:"auto;pk"`
	Content   string       `orm:"type(text)"`
	Tutorial  *TutorialORM `orm:"rel(fk);column(tutorial_id)"`
	CreatedAt time.Time    `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time    `orm:"auto_now;type(datetime)"`
	IsDeleted bool         `orm:"column(is_deleted);default(false)"`
}

func (c *CommentORM) TableName() string {
	return "comments"
}
