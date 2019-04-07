package services

import (
	"fmt"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/jaymist/greenretro/models"
)

// CreateUserAndTeam creates a user and their corresponding team.
func CreateUserAndTeam(tx *pop.Connection, u *models.User) (*validate.Errors, error) {
	t := &models.Team{
		Name: u.FirstName,
	}

	verrs, err := tx.ValidateAndCreate(t)
	if err != nil || verrs.HasAny() {
		return verrs, err
	}

	fmt.Printf("Creating user.")
	u.Teams = append(u.Teams, *t)
	verrs, err = u.Create(tx)
	return verrs, err
}
