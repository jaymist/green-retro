package services

import (
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/jaymist/greenretro/models"
)

// CreateUserAndTeam creates a user and their corresponding team.
func CreateUserAndTeam(tx *pop.Connection, u *models.User) (*validate.Errors, error) {
	verrs, err := u.Create(tx)
	if err != nil || verrs.HasAny() {
		return verrs, err
	}

	t := &models.Team{
		Name:  u.FirstName,
		Users: models.Users{*u},
	}

	verrs, err = tx.ValidateAndCreate(t)
	return verrs, err
}
