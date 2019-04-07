package services_test

import (
	"testing"

	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/suite"
)

type ServiceSuite struct {
	*suite.Model
}

func Test_ServiceSuite(t *testing.T) {
	model, err := suite.NewModelWithFixtures(packr.NewBox("../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	ss := &ServiceSuite{
		Model: model,
	}
	suite.Run(t, ss)
}
