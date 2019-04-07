package services_test

import (
	"github.com/jaymist/greenretro/models"
	"github.com/jaymist/greenretro/services"
)

func (ss *ServiceSuite) Test_User_Team_Create() {
	count, err := ss.DB.Count("teams")
	ss.NoError(err)
	ss.Equal(0, count)

	count, err = ss.DB.Count("users")
	ss.NoError(err)
	ss.Equal(0, count)

	u := &models.User{
		FirstName: "Bugs",
		LastName:  "Bunny",
		Email:     "bugs@acme.com",
		Password:  "password",
	}

	verrs, err := services.CreateUserAndTeam(ss.DB, u)
	ss.NoError(err)
	ss.False(verrs.HasAny())

	count, err = ss.DB.Count("teams")
	ss.NoError(err)
	ss.Equal(1, count)

	count, err = ss.DB.Count("users")
	ss.NoError(err)
	ss.Equal(1, count)

	count, err = ss.DB.Count("users_teams")
	ss.NoError(err)
	ss.Equal(1, count)
}

func (ss *ServiceSuite) Test_User_Team_Create_With_Empty_Name() {
	count, err := ss.DB.Count("teams")
	ss.NoError(err)
	ss.Equal(0, count)

	count, err = ss.DB.Count("users")
	ss.NoError(err)
	ss.Equal(0, count)

	u := &models.User{
		LastName: "Bunny",
		Email:    "bugs@acme.com",
		Password: "password",
	}
	ss.Zero(u.PasswordHash)

	verrs, err := services.CreateUserAndTeam(ss.DB, u)
	ss.NoError(err)
	ss.True(verrs.HasAny())
	ss.Equal(1, len(verrs.Keys()))
	ss.Equal("first_name", verrs.Keys()[0])
	ss.EqualError(verrs, "Name can not be blank.")

	count, err = ss.DB.Count("teams")
	ss.NoError(err)
	ss.Equal(0, count)

	count, err = ss.DB.Count("users")
	ss.NoError(err)
	ss.Equal(0, count)

	count, err = ss.DB.Count("users_teams")
	ss.NoError(err)
	ss.Equal(0, count)
}
