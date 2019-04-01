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
	as.LoadFixture("user accounts")

	res := as.HTML("/signin").Post(models.User{
		Email:    "bugs@acme.com",
		Password: "password",
	})
	as.Equal(302, res.Code)
	as.Equal("/", res.Location())

	res = as.HTML(res.Location()).Get()
	as.Contains(res.Body.String(), "bugs's Team")
	as.Contains(res.Body.String(), "Sign Out")

	as.Session.Clear()
	res = as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}
