package controllers

import (
	"encoding/json"
	"net/http"

	"oati-crud-comentarios/dtos"

	beego "github.com/beego/beego/v2/server/web"
)

type CommentController struct {
	beego.Controller
}

// GetByTutorialId
// @Title Listar comentarios de un tutorial
// @Description Obtiene todos los comentarios de un tutorial
// @Param tutorialId path int true "ID del tutorial"
// @Success 200 {object} dtos.CommentListResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @router /tutorials/:tutorialId/comments [get]
func (c *CommentController) GetByTutorialId() {
	tutorialId, err := parseID(c.Ctx.Input.Param(":tutorialId"))
	if err != nil {
		respondError(&c.Controller, http.StatusBadRequest, err.Error())
		return
	}

	comments, err := commentService.GetByTutorialId(tutorialId)
	if err != nil {
		status, msg := mapServiceError(err)
		respondError(&c.Controller, status, msg)
		return
	}
	respondJSON(&c.Controller, http.StatusOK, toCommentListResponse(comments))
}

// Create
// @Title Crear comentario
// @Description Crea un comentario asociado a un tutorial
// @Param tutorialId path int true "ID del tutorial"
// @Param body body dtos.CreateCommentRequest true "Datos del comentario"
// @Success 201 {object} dtos.CommentResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @router /tutorials/:tutorialId/comments [post]
func (c *CommentController) Create() {
	tutorialId, err := parseID(c.Ctx.Input.Param(":tutorialId"))
	if err != nil {
		respondError(&c.Controller, http.StatusBadRequest, err.Error())
		return
	}

	var req dtos.CreateCommentRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		respondError(&c.Controller, http.StatusBadRequest, "json inválido")
		return
	}

	comment, err := commentService.Create(req.Content, tutorialId)
	if err != nil {
		status, msg := mapServiceError(err)
		respondError(&c.Controller, status, msg)
		return
	}
	respondJSON(&c.Controller, http.StatusCreated, toCommentResponse(comment))
}

// Update
// @Title Actualizar comentario
// @Description Actualiza el contenido de un comentario
// @Param id path int true "ID del comentario"
// @Param body body dtos.UpdateCommentRequest true "Datos del comentario"
// @Success 200 {object} dtos.CommentResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @router /comments/:id [put]
func (c *CommentController) Update() {
	id, err := parseID(c.Ctx.Input.Param(":id"))
	if err != nil {
		respondError(&c.Controller, http.StatusBadRequest, err.Error())
		return
	}

	var req dtos.UpdateCommentRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		respondError(&c.Controller, http.StatusBadRequest, "json inválido")
		return
	}

	comment, err := commentService.Update(id, req.Content)
	if err != nil {
		status, msg := mapServiceError(err)
		respondError(&c.Controller, status, msg)
		return
	}
	respondJSON(&c.Controller, http.StatusOK, toCommentResponse(comment))
}

// Delete
// @Title Eliminar comentario
// @Description Elimina un comentario por ID
// @Param id path int true "ID del comentario"
// @Success 204 "No Content"
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @router /comments/:id [delete]
func (c *CommentController) Delete() {
	id, err := parseID(c.Ctx.Input.Param(":id"))
	if err != nil {
		respondError(&c.Controller, http.StatusBadRequest, err.Error())
		return
	}

	if err := commentService.Delete(id); err != nil {
		status, msg := mapServiceError(err)
		respondError(&c.Controller, status, msg)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
}
