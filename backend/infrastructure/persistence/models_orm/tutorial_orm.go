package models_orm

import "time"

type TutorialORM struct {
	Id          int            `orm:"auto;pk"`
	Title       string         `orm:"size(200)"`
	Description string         `orm:"type(text)"`
	PublishedAt time.Time      `orm:"type(datetime)"`
	CreatedAt   time.Time      `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time      `orm:"auto_now;type(datetime)"`
	IsDeleted   bool           `orm:"column(is_deleted);default(false)"`
	Comments    []*CommentORM  `orm:"reverse(many)"`
}

func (t *TutorialORM) TableName() string {
	return "tutorials"
}
