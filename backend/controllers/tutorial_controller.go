package controllers

import (
	"encoding/json"
	"net/http"

	"oati-crud-comentarios/dtos"

	beego "github.com/beego/beego/v2/server/web"
)

type TutorialController struct {
	beego.Controller
}

// GetAll
// @Title Listar tutoriales
// @Description Obtiene todos los tutoriales sin comentarios anidados
// @Success 200 {object} dtos.TutorialListResponse
// @Failure 500 {object} dtos.ErrorResponse
// @router /tutorials [get]
func (c *TutorialController) GetAll() {
	tutorials, err := tutorialService.GetAll()
	if err != nil {
		status, msg := mapServiceError(err)
		respondError(&c.Controller, status, msg)
		return
	}
	respondJSON(&c.Controller, http.StatusOK, toTutorialListResponse(tutorials))
}

// GetById
// @Title Obtener tutorial
// @Description Obtiene un tutorial por ID con sus comentarios anidados
// @Param id path int true "ID del tutorial"
// @Success 200 {object} dtos.TutorialDetailResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @router /tutorials/:id [get]
func (c *TutorialController) GetById() {
	id, err := parseID(c.Ctx.Input.Param(":id"))
	if err != nil {
		respondError(&c.Controller, http.StatusBadRequest, err.Error())
		return
	}

	tutorial, err := tutorialService.GetById(id)
	if err != nil {
		status, msg := mapServiceError(err)
		respondError(&c.Controller, status, msg)
		return
	}
	respondJSON(&c.Controller, http.StatusOK, toTutorialDetailResponse(tutorial))
}

// Create
// @Title Crear tutorial
// @Description Crea un nuevo tutorial. published_at debe estar en formato RFC3339 (ej: 2026-06-18T00:00:00Z)
// @Param body body dtos.CreateTutorialRequest true "Datos del tutorial"
// @Success 201 {object} dtos.TutorialResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @router /tutorials [post]
func (c *TutorialController) Create() {
	var req dtos.CreateTutorialRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		respondError(&c.Controller, http.StatusBadRequest, "json inválido")
		return
	}

	publishedAt, err := parsePublishedAt(req.PublishedAt)
	if err != nil {
		respondError(&c.Controller, http.StatusBadRequest, err.Error())
		return
	}

	tutorial, err := tutorialService.Create(req.Title, req.Description, publishedAt)
	if err != nil {
		status, msg := mapServiceError(err)
		respondError(&c.Controller, status, msg)
		return
	}
	respondJSON(&c.Controller, http.StatusCreated, toTutorialResponse(tutorial))
}

// Update
// @Title Actualizar tutorial
// @Description Actualiza un tutorial existente
// @Param id path int true "ID del tutorial"
// @Param body body dtos.UpdateTutorialRequest true "Datos del tutorial"
// @Success 200 {object} dtos.TutorialResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @router /tutorials/:id [put]
func (c *TutorialController) Update() {
	id, err := parseID(c.Ctx.Input.Param(":id"))
	if err != nil {
		respondError(&c.Controller, http.StatusBadRequest, err.Error())
		return
	}

	var req dtos.UpdateTutorialRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		respondError(&c.Controller, http.StatusBadRequest, "json inválido")
		return
	}

	publishedAt, err := parsePublishedAt(req.PublishedAt)
	if err != nil {
		respondError(&c.Controller, http.StatusBadRequest, err.Error())
		return
	}

	tutorial, err := tutorialService.Update(id, req.Title, req.Description, publishedAt)
	if err != nil {
		status, msg := mapServiceError(err)
		respondError(&c.Controller, status, msg)
		return
	}
	respondJSON(&c.Controller, http.StatusOK, toTutorialResponse(tutorial))
}

// Delete
// @Title Eliminar tutorial
// @Description Elimina un tutorial y sus comentarios asociados
// @Param id path int true "ID del tutorial"
// @Success 204 "No Content"
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @router /tutorials/:id [delete]
func (c *TutorialController) Delete() {
	id, err := parseID(c.Ctx.Input.Param(":id"))
	if err != nil {
		respondError(&c.Controller, http.StatusBadRequest, err.Error())
		return
	}

	if err := tutorialService.Delete(id); err != nil {
		status, msg := mapServiceError(err)
		respondError(&c.Controller, status, msg)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
}
