package grifts

import (
	"github.com/gobuffalo/pop"
	"github.com/jaymist/greenretro/models"
	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		return models.DB.Transaction(func(tx *pop.Connection) error {
			return createModels(tx)
		})
	})
})

func createModels(tx *pop.Connection) error {
	if err := tx.TruncateAll(); err != nil {
		return errors.WithStack(err)
	}

	// Add DB seeding stuff here
	u := &models.User{
		Email:                "bugs@acme.com",
		FirstName:            "Bugs",
		LastName:             "Bunny",
		Password:             "password",
		PasswordConfirmation: "password",
	}

	if _, err := u.Create(tx); err != nil {
		return errors.WithStack(err)
	}

	t := &models.Team{
		Name:  "Bugs's Team",
		Users: models.Users{*u},
	}

	if err := tx.Create(t); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
