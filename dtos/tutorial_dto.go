package dtos

import "time"

type CreateTutorialRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishedAt string `json:"published_at"`
}

type UpdateTutorialRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishedAt string `json:"published_at"`
}

type TutorialResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TutorialListResponse struct {
	Data  []*TutorialResponse `json:"data"`
	Total int                 `json:"total"`
}

type TutorialDetailResponse struct {
	Id          int                `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	PublishedAt time.Time          `json:"published_at"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Comments    []*CommentResponse `json:"comments"`
}
