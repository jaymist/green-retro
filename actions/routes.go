package actions

import (
	"github.com/gobuffalo/buffalo"
)

// define application routes here.
func routes(app *buffalo.App) {
	// FRONTEND routes and middleware
	frontendRoutes(app)

	// API routes
	apiRoutes(app)

	app.ServeFiles("/", assetsBox) // serve files from the public directory
}

func frontendRoutes(app *buffalo.App) {
	f := app.Group("/")

	f.Use(SetCurrentUser)
	f.Use(Authorize)

	f.GET("/", HomeHandler)
	f.GET("/register", UsersNew)
	f.POST("/users", UsersCreate)
	f.GET("/signin", AuthNew)
	f.POST("/signin", AuthCreate)
	f.DELETE("/signout", AuthDestroy)

	f.Middleware.Skip(Authorize, HomeHandler, UsersNew, UsersCreate, AuthNew, AuthCreate)
}

func apiRoutes(app *buffalo.App) {
	a := app.Group("/api")

	a.Use(BasicAuthorize)
}
