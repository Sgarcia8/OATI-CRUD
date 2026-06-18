package main

import (
	"log"

	"oati-crud-comentarios/infraestructure/database"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	beego.Run()
}
