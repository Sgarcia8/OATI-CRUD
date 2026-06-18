package routers

import (
	"oati-crud-comentarios/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/api/v1/tutorials", &controllers.TutorialController{}, "get:GetAll;post:Create")
	beego.Router("/api/v1/tutorials/:id", &controllers.TutorialController{}, "get:GetById;put:Update;delete:Delete")

	beego.Router("/api/v1/tutorials/:tutorialId/comments", &controllers.CommentController{}, "get:GetByTutorialId;post:Create")
	beego.Router("/api/v1/comments/:id", &controllers.CommentController{}, "put:Update;delete:Delete")
}
