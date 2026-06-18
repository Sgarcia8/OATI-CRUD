package models

import "time"

type Tutorial struct {
	Id          int    `orm:"auto"`
	Title       string `orm:"size(200)"`
	Description string `orm:"type(text)"`
	PublishedAt time.Time
	CreatedAt   time.Time `orm:"auto_now_add"`
}
