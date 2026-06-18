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
