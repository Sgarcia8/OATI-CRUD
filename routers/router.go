// @APIVersion 1.0.0
// @Title OATI CRUD API
// @Description API REST para gestión de tutoriales y comentarios
// @Contact sjgarciat8@gmail.com
// @TermsOfServiceUrl http://swagger.io/terms/
package routers

import (
	"oati-crud-comentarios/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSInclude(
			&controllers.TutorialController{},
			&controllers.CommentController{},
		),
	)
	beego.AddNamespace(ns)
}
