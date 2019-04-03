package actions

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/jaymist/greenretro/models"
	"github.com/pkg/errors"
)

// UsersNew renders the registration page.
func UsersNew(c buffalo.Context) error {
	u := models.User{}
	c.Set("user", u)
	return c.Render(200, r.HTML("users/new.html"))
}

// UsersCreate registers a new user with the application.
func UsersCreate(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := u.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("user", u)
		c.Set("errors", verrs)
		return c.Render(200, r.HTML("users/new.html"))
	}

	c.Session().Set("current_user_id", u.ID)
	c.Flash().Add("success", "Welcome to Buffalo!")

	return c.Redirect(302, "/")
}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				return errors.WithStack(err)
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Session().Set("redirectURL", c.Request().URL.String())

			err := c.Session().Save()
			if err != nil {
				return errors.WithStack(err)
			}

			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}

// BasicAuthorize adds basic auth support to requests.
func BasicAuthorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="Basic Authentication"`)
			return c.Error(http.StatusUnauthorized, errors.New("Unauthorized"))
		}
		c.Logger().WithField("authorization", authHeader).Info("Header")

		fields := strings.Split(authHeader, " ")
		if len(fields) < 2 {
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="Basic Authentication"`)
			return c.Error(http.StatusUnauthorized, errors.New("Unauthorized"))
		}
		token := fields[1]
		c.Logger().WithField("token", token).Info("Token")

		b, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="Basic Authentication"`)
			return c.Error(http.StatusUnauthorized, errors.New("Unauthorized"))
		}

		pair := strings.SplitN(string(b), ":", 2)
		params := map[string]string{
			"user":     pair[0],
			"password": pair[1],
		}

		c.Logger().WithField("user details", params).Info("User details")
		return next(c)
	}
}
