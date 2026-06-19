package main

import (
	"log"

	"oati-crud-comentarios/controllers"
	"oati-crud-comentarios/infrastructure/database"
	"oati-crud-comentarios/infrastructure/persistence/models_orm"
	"oati-crud-comentarios/infrastructure/persistence/repository_impl"
	"oati-crud-comentarios/middleware"
	"oati-crud-comentarios/services"
	_ "oati-crud-comentarios/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	models_orm.RegisterModels()
	orm.RunSyncdb("default", false, true)

	tutorialRepo := repository_impl.NewTutorialRepository()
	commentRepo := repository_impl.NewCommentRepository()
	tutorialSvc := services.NewTutorialService(tutorialRepo)
	commentSvc := services.NewCommentService(commentRepo, tutorialRepo)
	controllers.Init(tutorialSvc, commentSvc)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, middleware.CorsFilter, beego.WithReturnOnOutput(true))

	beego.Run()
}
