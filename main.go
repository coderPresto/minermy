package main

import (
  "minermy/web/controllers"

  "github.com/kataras/iris"
  "github.com/kataras/iris/mvc"
)

func main() {
  app := iris.New()
  app.Logger().SetLevel("debug")

  mvc.New(app.Party("/user")).Handle(new(controllers.UserController))
  //mvc.Configure(app.Party("/user"), user)

  app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
  )
}
