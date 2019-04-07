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
