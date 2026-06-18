package domain

import "time"

type Tutorial struct {
	Id          int
	Title       string
	Description string
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool
	Comments    []*Comment
}

func NewTutorial(title, description string, publishedAt time.Time) *Tutorial {
	now := time.Now()
	return &Tutorial{
		Title:       title,
		Description: description,
		PublishedAt: publishedAt,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (t *Tutorial) Update(title, description string, publishedAt time.Time) {
	t.Title = title
	t.Description = description
	t.PublishedAt = publishedAt
	t.UpdatedAt = time.Now()
}

func (t *Tutorial) IsValid() bool {
	return t.Title != "" && t.Description != ""
}

func (t *Tutorial) AddComment(content string) *Comment {
	comment := NewComment(content, t.Id)
	t.Comments = append(t.Comments, comment)
	return comment
}

func (t *Tutorial) SoftDelete() {
	t.IsDeleted = true
	t.UpdatedAt = time.Now()
}
