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

func (ms *ModelSuite) Test_Team_Create_Missing_Name() {
	count, err := ms.DB.Count("teams")
	ms.NoError(err)
	ms.Equal(0, count)

	t := &models.Team{}

	verrs, err := t.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())
	ms.Equal(1, len(verrs.Keys()))
	ms.Equal("name", verrs.Keys()[0])
	ms.EqualError(verrs, "Name can not be blank.")

	count, err = ms.DB.Count("teams")
	ms.NoError(err)
	ms.Equal(0, count)
}
