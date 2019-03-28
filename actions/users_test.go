package actions_test

import (
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
		Email:                "mark@example.com",
		FirstName:            "Mark",
		LastName:             "Example",
		Password:             "password",
		PasswordConfirmation: "password",
	}

	res := as.HTML("/users").Post(u)
	as.Equal(302, res.Code)

	count, err = as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)
}
