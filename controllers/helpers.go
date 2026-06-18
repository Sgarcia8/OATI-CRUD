package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"oati-crud-comentarios/domain"
	"oati-crud-comentarios/dtos"
	"oati-crud-comentarios/services"

	beego "github.com/beego/beego/v2/server/web"
)

func parseID(param string) (int, error) {
	id, err := strconv.Atoi(param)
	if err != nil || id <= 0 {
		return 0, errors.New("id inválido")
	}
	return id, nil
}

func parsePublishedAt(s string) (time.Time, error) {
	if s == "" {
		return time.Time{}, errors.New("published_at es obligatorio")
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}, errors.New("published_at debe estar en formato RFC3339")
	}
	return t, nil
}

func respondJSON(c *beego.Controller, status int, payload interface{}) {
	c.Ctx.ResponseWriter.WriteHeader(status)
	c.Data["json"] = payload
	c.ServeJSON()
}

func respondError(c *beego.Controller, status int, message string) {
	respondJSON(c, status, dtos.NewErrorResponse(status, message))
}

func mapServiceError(err error) (int, string) {
	switch {
	case errors.Is(err, services.ErrTutorialNotFound),
		errors.Is(err, services.ErrCommentNotFound),
		errors.Is(err, services.ErrTutorialNotFoundForComment):
		return http.StatusNotFound, err.Error()
	case errors.Is(err, services.ErrTutorialInvalidData),
		errors.Is(err, services.ErrCommentInvalidData):
		return http.StatusBadRequest, err.Error()
	default:
		return http.StatusInternalServerError, "error interno del servidor"
	}
}

func toTutorialResponse(t *domain.Tutorial) *dtos.TutorialResponse {
	return &dtos.TutorialResponse{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		PublishedAt: t.PublishedAt,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func toTutorialListResponse(tutorials []*domain.Tutorial) *dtos.TutorialListResponse {
	data := make([]*dtos.TutorialResponse, len(tutorials))
	for i, t := range tutorials {
		data[i] = toTutorialResponse(t)
	}
	return &dtos.TutorialListResponse{
		Data:  data,
		Total: len(data),
	}
}

func toTutorialDetailResponse(t *domain.Tutorial) *dtos.TutorialDetailResponse {
	comments := make([]*dtos.CommentResponse, len(t.Comments))
	for i, c := range t.Comments {
		comments[i] = toCommentResponse(c)
	}
	return &dtos.TutorialDetailResponse{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		PublishedAt: t.PublishedAt,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		Comments:    comments,
	}
}

func toCommentResponse(c *domain.Comment) *dtos.CommentResponse {
	return &dtos.CommentResponse{
		Id:         c.Id,
		Content:    c.Content,
		TutorialId: c.TutorialId,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}

func toCommentListResponse(comments []*domain.Comment) *dtos.CommentListResponse {
	data := make([]*dtos.CommentResponse, len(comments))
	for i, c := range comments {
		data[i] = toCommentResponse(c)
	}
	return &dtos.CommentListResponse{
		Data:  data,
		Total: len(data),
	}
}
