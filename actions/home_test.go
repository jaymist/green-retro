package actions_test

import "github.com/jaymist/greenretro/models"

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/").Get()
	as.Equal(200, res.Code)

	body := res.Body.String()
	as.Contains(body, "Sign In")
	as.Contains(body, "Register")

	as.Contains(body, "/register")
	as.Contains(body, "/signin")
}

func (as *ActionSuite) Test_HomeHandler_LoggedIn() {
	u := &models.User{
		Email:                "mark@example.com",
		FirstName:            "Mark",
		LastName:             "Example",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())
	as.Session.Set("current_user_id", u.ID)

	res := as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign Out")

	as.Session.Clear()
	res = as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}
