package controllers

import "github.com/kataras/iris"

type AdminController struct {
	Ctx iris.Context
}

func (c *AdminController) Get() {
	c.Ctx.ViewData("a", "b")
	c.Ctx.View("admin/index.html")
}
