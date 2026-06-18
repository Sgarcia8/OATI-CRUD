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

func (c *TutorialController) GetAll() {
	tutorials, err := tutorialService.GetAll()
	if err != nil {
		status, msg := mapServiceError(err)
		respondError(&c.Controller, status, msg)
		return
	}
	respondJSON(&c.Controller, http.StatusOK, toTutorialListResponse(tutorials))
}

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
