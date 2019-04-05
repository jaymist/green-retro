package models_test

import (
	"github.com/jaymist/greenretro/models"
)

func (ms *ModelSuite) Test_Team_Create() {
	count, err := ms.DB.Count("teams")
	ms.NoError(err)
	ms.Equal(0, count)

	t := &models.Team{
		Name: "Mark's Team",
	}

	verrs, err := t.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	count, err = ms.DB.Count("teams")
	ms.NoError(err)
	ms.Equal(1, count)
}
