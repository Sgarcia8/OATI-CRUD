package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["oati-crud-comentarios/controllers:CommentController"] = append(beego.GlobalControllerRouter["oati-crud-comentarios/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/comments/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oati-crud-comentarios/controllers:CommentController"] = append(beego.GlobalControllerRouter["oati-crud-comentarios/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/comments/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oati-crud-comentarios/controllers:CommentController"] = append(beego.GlobalControllerRouter["oati-crud-comentarios/controllers:CommentController"],
        beego.ControllerComments{
            Method: "GetByTutorialId",
            Router: `/tutorials/:tutorialId/comments`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oati-crud-comentarios/controllers:CommentController"] = append(beego.GlobalControllerRouter["oati-crud-comentarios/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/tutorials/:tutorialId/comments`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oati-crud-comentarios/controllers:TutorialController"] = append(beego.GlobalControllerRouter["oati-crud-comentarios/controllers:TutorialController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/tutorials`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oati-crud-comentarios/controllers:TutorialController"] = append(beego.GlobalControllerRouter["oati-crud-comentarios/controllers:TutorialController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/tutorials`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oati-crud-comentarios/controllers:TutorialController"] = append(beego.GlobalControllerRouter["oati-crud-comentarios/controllers:TutorialController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/tutorials/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oati-crud-comentarios/controllers:TutorialController"] = append(beego.GlobalControllerRouter["oati-crud-comentarios/controllers:TutorialController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/tutorials/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oati-crud-comentarios/controllers:TutorialController"] = append(beego.GlobalControllerRouter["oati-crud-comentarios/controllers:TutorialController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/tutorials/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
