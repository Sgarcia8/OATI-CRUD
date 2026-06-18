package dtos

import "time"

type CreateCommentRequest struct {
	Content string `json:"content"`
}

type UpdateCommentRequest struct {
	Content string `json:"content"`
}

type CommentResponse struct {
	Id         int       `json:"id"`
	Content    string    `json:"content"`
	TutorialId int       `json:"tutorial_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CommentListResponse struct {
	Data  []*CommentResponse `json:"data"`
	Total int                `json:"total"`
}
