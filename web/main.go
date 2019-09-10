package main

/**

 */
import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"xslt-platform/web/controllers"
)

func newApp() *iris.Application {
	app := iris.Default()
	//app.Logger().SetLevel("debug")
	tmpl := iris.HTML("./views", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(tmpl)
	app.Any("/test", func(c context.Context) {
		c.WriteString("test")
	})
	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})
	mvc.New(app.Party("/admin")).Handle(new(controllers.AdminController))
	mvc.New(app).Handle(new(controllers.HomeController))

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("shared/error.html")
	})
	app.StaticWeb("/public", "./public")
	return app
}

func main() {
	app := newApp()

	app.Run(
		iris.Addr(":8080"),
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}
