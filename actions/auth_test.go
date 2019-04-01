package actions_test

import (
	"github.com/jaymist/greenretro/models"
)

func (as *ActionSuite) Test_Auth_New() {
	res := as.HTML("/signin").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}

func (as *ActionSuite) Test_Auth_Create() {
	as.LoadFixture("user accounts")

	res := as.HTML("/signin").Post(models.User{
		Email:    "bugs@acme.com",
		Password: "password",
	})
	as.Equal(302, res.Code)
	as.Equal("/", res.Location())
}

func (as *ActionSuite) Test_Auth_Create_Redirect() {
	as.LoadFixture("user accounts")
	as.Session.Set("redirectURL", "/some/url")

	res := as.HTML("/signin").Post(models.User{
		Email:    "bugs@acme.com",
		Password: "password",
	})
	as.Equal(302, res.Code)
	as.Equal(res.Location(), "/some/url")
}

func (as *ActionSuite) Test_Auth_Create_UnknownUser() {
	res := as.HTML("/signin").Post(models.User{
		Email:    "mark@example.com",
		Password: "password",
	})
	as.Equal(422, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}

func (as *ActionSuite) Test_Auth_Create_BadPassword() {
	as.LoadFixture("user accounts")

	res := as.HTML("/signin").Post(models.User{
		Email:    "bugs@acme.com",
		Password: "bad_password",
	})
	as.Equal(422, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}
