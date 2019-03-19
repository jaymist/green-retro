package actions

import (
	"fmt"

	"github.com/jaymist/greenretro/models"
)

func (as *ActionSuite) Test_CompaniesResource_List() {
	// as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_CompaniesResource_Show() {
	// as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_CompaniesResource_Create() {
	company := &models.Company{Name: "Test Company"}
	result := as.HTML("/companies").Post(company)

	// assert that the response status code was 302 as.Equal(302, res.Code)
	// retreive the first Widget from the database
	err := as.DB.First(company)
	as.NoError(err)
	as.NotZero(company.ID)
	// assert the Widget title was saved correctly
	as.Equal("Test Company", company.Name)
	// assert the redirect was sent to the place expected
	as.Equal(fmt.Sprintf("/companies/%s", company.ID), result.Location())
}

func (as *ActionSuite) Test_CompaniesResource_Update() {
	// as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_CompaniesResource_Destroy() {
	// as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_CompaniesResource_New() {
	// as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_CompaniesResource_Edit() {
	// as.Fail("Not Implemented!")
}
