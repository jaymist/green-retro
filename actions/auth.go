package actions

import (
	"database/sql"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/jaymist/greenretro/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// AuthNew loads the signin page
func AuthNew(c buffalo.Context) error {
	c.Set("user", models.User{})
	return c.Render(200, r.HTML("auth/new.html"))
}

// AuthCreate attempts to log the user in with an existing account.
func AuthCreate(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	if err := authenticate(c, u.Email, u.Password); err != nil {
		return c.Error(http.StatusUnauthorized, err)
	}

	c.Session().Set("current_user_id", u.ID)
	c.Flash().Add("success", "Welcome Back to Green Retro!")

	redirectURL := "/"
	if redir, ok := c.Session().Get("redirectURL").(string); ok {
		redirectURL = redir
	}

	return c.Redirect(http.StatusFound, redirectURL)
}

// AuthDestroy clears the session and logs a user out
func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("success", "You have been logged out!")
	return c.Redirect(302, "/")
}

// BasicAuth adds basic auth support to requests.
func BasicAuth(next buffalo.Handler) buffalo.Handler {
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

		params := strings.SplitN(string(b), ":", 2)
		if err := authenticate(c, params[0], params[1]); err != nil {
			return c.Error(http.StatusUnauthorized, err)
		}

		return next(c)
	}
}

func authenticate(c buffalo.Context, email, password string) error {
	tx := c.Value("tx").(*pop.Connection)
	u := &models.User{}

	// find a user with the email
	err := tx.Where("email = ?", strings.ToLower(strings.TrimSpace(email))).First(u)

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with the supplied email address.
			return errors.New("invalid credentials")
		}
		return errors.WithStack(err)
	}

	// confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		return errors.New("invalid credentials")
	}

	return nil
}
