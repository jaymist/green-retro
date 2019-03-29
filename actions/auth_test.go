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

	u := &models.User{
		Email:    "bugs@acme.com",
		Password: "password",
	}

	res := as.HTML("/signin").Post(u)
	as.Equal(302, res.Code)
	as.Equal("/", res.Location())
}

func (as *ActionSuite) Test_Auth_Create_Redirect() {
	as.LoadFixture("user accounts")

	u := &models.User{
		Email:    "bugs@acme.com",
		Password: "password",
	}

	as.Session.Set("redirectURL", "/some/url")

	res := as.HTML("/signin").Post(u)
	as.Equal(302, res.Code)
	as.Equal(res.Location(), "/some/url")
}

func (as *ActionSuite) Test_Auth_Create_UnknownUser() {
	u := &models.User{
		Email:    "mark@example.com",
		Password: "password",
	}
	res := as.HTML("/signin").Post(u)
	as.Equal(422, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}

func (as *ActionSuite) Test_Auth_Create_BadPassword() {
	as.LoadFixture("user accounts")

	u := &models.User{
		Email:    "bugs@acme.com",
		Password: "bad",
	}

	res := as.HTML("/signin").Post(u)
	as.Equal(422, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}
