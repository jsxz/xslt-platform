package controllers

import (
	"github.com/kataras/iris/mvc"
)

type HomeController struct {
}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *HomeController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}
