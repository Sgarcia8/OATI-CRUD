package main

import (
	_ "oati-crud-comentarios/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

