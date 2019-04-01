package actions

import (
	"github.com/gobuffalo/buffalo"
)

// define application routes here.
func routes(app *buffalo.App) {
	app.GET("/", HomeHandler)

	app.Use(SetCurrentUser)
	app.Use(Authorize)

	app.GET("/register", UsersNew)
	app.POST("/users", UsersCreate)
	app.GET("/signin", AuthNew)
	app.POST("/signin", AuthCreate)
	app.DELETE("/signout", AuthDestroy)

	// API Endpoints
	app.Resource("/teams", TeamsResource{})

	app.Middleware.Skip(Authorize, HomeHandler, UsersNew, UsersCreate, AuthNew, AuthCreate)
	app.ServeFiles("/", assetsBox) // serve files from the public directory
}
