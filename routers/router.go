package routers

import (
	"oati-crud-comentarios/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.CommentController{})
}
