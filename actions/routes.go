package actions

import (
	"github.com/gobuffalo/buffalo"
)

// define application routes here.
func routes(app *buffalo.App) {
	uiRoutes(app)
	apiRoutes(app)

	app.ServeFiles("/", assetsBox) // serve files from the public directory
}

func uiRoutes(app *buffalo.App) {
	ui := app.Group("/")

	ui.Use(SetCurrentUser)
	ui.Use(Authorize)

	ui.GET("/", HomeHandler)
	ui.GET("/register", UsersNew)
	ui.POST("/users", UsersCreate)
	ui.GET("/signin", AuthNew)
	ui.POST("/signin", AuthCreate)
	ui.DELETE("/signout", AuthDestroy)
	ui.Middleware.Skip(Authorize, HomeHandler, UsersNew, UsersCreate, AuthNew, AuthCreate)
}

func apiRoutes(app *buffalo.App) {
	api := app.Group("/api")

	api.Use(Authorize)
}
