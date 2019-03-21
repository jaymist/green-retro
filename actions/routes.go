package actions

import (
	"github.com/gobuffalo/buffalo"
)

// Routes defines all routes for the application
func Routes(app *buffalo.App) {
	app.GET("/", HomeHandler)

	app.Use(SetCurrentUser)
	app.Use(Authorize)
	app.GET("/register", Register)
	app.POST("/users", UsersCreate)
	app.GET("/signin", AuthNew)
	app.POST("/signin", AuthCreate)
	app.DELETE("/signout", AuthDestroy)
	app.Middleware.Skip(Authorize, HomeHandler, Register, UsersCreate, AuthNew, AuthCreate)

	app.ServeFiles("/", assetsBox) // serve files from the public directory
}
