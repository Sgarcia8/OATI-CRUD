package domain

import "time"

type Comment struct {
	Id         int
	Content    string
	TutorialId int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsDeleted  bool
}

func NewComment(content string, tutorialId int) *Comment {
	now := time.Now()
	return &Comment{
		Content:    content,
		TutorialId: tutorialId,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func (c *Comment) Update(content string) {
	c.Content = content
	c.UpdatedAt = time.Now()
}

func (c *Comment) IsValid() bool {
	return c.Content != "" && c.TutorialId > 0
}

func (c *Comment) SoftDelete() {
	c.IsDeleted = true
	c.UpdatedAt = time.Now()
}
