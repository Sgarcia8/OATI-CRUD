package dtos

import "time"

type CreateTutorialRequest struct {
	Title       string `json:"title" example:"Tutorial de prueba" description:"Título del tutorial"`
	Description string `json:"description" example:"Descripción del tutorial" description:"Descripción del tutorial"`
	PublishedAt string `json:"published_at" example:"2026-06-18T00:00:00Z" description:"Fecha de publicación del tutorial"`
}

type UpdateTutorialRequest struct {
	Title       string `json:"title" example:"Tutorial de prueba" description:"Título del tutorial"`
	Description string `json:"description" example:"Descripción del tutorial" description:"Descripción del tutorial"`
	PublishedAt string `json:"published_at" example:"2026-06-18T00:00:00Z" description:"Fecha de publicación del tutorial"`
}

type TutorialResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title" example:"Tutorial de prueba" description:"Título del tutorial"`
	Description string    `json:"description" example:"Descripción del tutorial" description:"Descripción del tutorial"`
	PublishedAt time.Time `json:"published_at" example:"2026-06-18T00:00:00Z" description:"Fecha de publicación del tutorial"`
	CreatedAt   time.Time `json:"created_at" example:"2026-06-18T00:00:00Z" description:"Fecha de creación del tutorial"`
	UpdatedAt   time.Time `json:"updated_at" example:"2026-06-18T00:00:00Z" description:"Fecha de actualización del tutorial"`
}

type TutorialListResponse struct {
	Data  []*TutorialResponse `json:"data"`
	Total int                 `json:"total"`
}

type TutorialDetailResponse struct {
	Id          int                `json:"id"`
	Title       string             `json:"title" example:"Tutorial de prueba" description:"Título del tutorial"`
	Description string             `json:"description" example:"Descripción del tutorial" description:"Descripción del tutorial"`
	PublishedAt time.Time          `json:"published_at" example:"2026-06-18T00:00:00Z" description:"Fecha de publicación del tutorial"`
	CreatedAt   time.Time          `json:"created_at" example:"2026-06-18T00:00:00Z" description:"Fecha de creación del tutorial"`
	UpdatedAt   time.Time          `json:"updated_at" example:"2026-06-18T00:00:00Z" description:"Fecha de actualización del tutorial"`
	Comments    []*CommentResponse `json:"comments" example:"[{'id': 1, 'content': 'Comentario de prueba', 'tutorial_id': 1, 'created_at': '2026-06-18T00:00:00Z', 'updated_at': '2026-06-18T00:00:00Z'}]"`
}
