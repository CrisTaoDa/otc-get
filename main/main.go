package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/view"
	"otc-get/app/controller"
	"otc-get/app/util"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(iris.Gzip)
	app.StaticWeb("/web", util.LocalPath()+"web")
	app.RegisterView(ViewHelper())
	mvc.Configure(app.Party("/").Layout("layout/layout.html"), Route)
	err := app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		fmt.Println("error", err)
	}

}


func ViewHelper() *view.HTMLEngine {
	// 解析html
	tpl := iris.HTML(util.LocalPath()+"web/html", ".html").Reload(false)
	return tpl
}

// 路由管理
func Route(app *mvc.Application) {
	app.Party("/").Handle(new(controller.IndexController))
}
