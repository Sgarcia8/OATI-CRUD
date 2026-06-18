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
	Content    string    `json:"content" example:"Comentario de prueba" description:"Contenido del comentario"`
	TutorialId int       `json:"tutorial_id" example:"1" description:"ID del tutorial"`
	CreatedAt  time.Time `json:"created_at" example:"2026-06-18T00:00:00Z" description:"Fecha de creación del comentario"`
	UpdatedAt  time.Time `json:"updated_at" example:"2026-06-18T00:00:00Z" description:"Fecha de actualización del comentario"`
}

type CommentListResponse struct {
	Data  []*CommentResponse `json:"data"`
	Total int                `json:"total"`
}
