package models_test

import (
	"github.com/jaymist/greenretro/models"
)

func (ms *ModelSuite) Test_User_Create() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Email:                "mark@example.com",
		FirstName:            "Mark",
		LastName:             "Example",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.PasswordHash)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) TestUserCreateWithMissingEmail() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		FirstName:            "Mark",
		LastName:             "Example",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())
	ms.T().Log("Validation errors: ", verrs)
	ms.Equal(1, len(verrs.Keys()))
	ms.Equal("email", verrs.Keys()[0])
	ms.EqualError(verrs, "Email can not be blank.")

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) TestUserCreateWithMissingPassword() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Email:     "mark@example.com",
		FirstName: "Mark",
		LastName:  "Example",
	}
	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())
	ms.Equal(1, len(verrs.Keys()))
	ms.Equal("password", verrs.Keys()[0])
	ms.EqualError(verrs, "Password can not be blank.")

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) TestUserCreateWithMissingFirstName() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Email:                "mark@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())
	ms.Equal(1, len(verrs.Keys()))
	ms.Equal("first_name", verrs.Keys()[0])
	ms.EqualError(verrs, "Name can not be blank.")

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_User_Create_ValidationErrors() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Password: "password",
	}
	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_User_Create_UserExists() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Email:                "mark@example.com",
		FirstName:            "Mark",
		LastName:             "Example",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.PasswordHash)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	u = &models.User{
		Email:    "mark@example.com",
		Password: "password",
	}
	verrs, err = u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}
