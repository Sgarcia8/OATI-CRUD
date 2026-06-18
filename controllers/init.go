package controllers

import "oati-crud-comentarios/services"

var (
	tutorialService *services.TutorialService
	commentService  *services.CommentService
)

func Init(ts *services.TutorialService, cs *services.CommentService) {
	tutorialService = ts
	commentService = cs
}
