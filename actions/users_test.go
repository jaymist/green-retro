package actions_test

import (
	"net/http"

	"github.com/jaymist/greenretro/models"
)

func (as *ActionSuite) Test_Users_New() {
	res := as.HTML("/register").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_Users_Create() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)

	u := &models.User{
		Email:     "mark@example.com",
		FirstName: "Mark",
		LastName:  "Example",
		Password:  "password",
	}

	res := as.HTML("/users").Post(u)
	as.Equal(302, res.Code)

	count, err = as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)

	count, err = as.DB.Count("teams")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.Eager().Last(u)
	as.NoError(err)

	as.NotEmpty(u.Teams)
	as.Equal(u.FirstName, u.Teams[0].Name)
}

func (as *ActionSuite) Test_Users_Create_Without_Email() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)

	u := &models.User{
		FirstName: "Mark",
		LastName:  "Example",
		Password:  "password",
	}

	res := as.HTML("/users").Post(u)
	as.Equal(http.StatusBadRequest, res.Code)

	count, err = as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)
}
