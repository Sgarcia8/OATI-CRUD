package controllers

import beego "github.com/beego/beego/v2/server/web"

type CommentController struct {
	beego.Controller
}

func (c *CommentController) Get() {
	c.Data["json"] = "Hello, World!"
	c.ServeJSON()
}
