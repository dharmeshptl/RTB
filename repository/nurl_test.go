package repository_test

import (
	"github.com/stretchr/testify/suite"
	"go_rtb/internal/model"
	"go_rtb/internal/repository"
	"go_rtb/internal/test"
	"go_rtb/internal/test/data_generator"
	"go_rtb/internal/tool/helper"
	"testing"
)

type NurlRepositoryTestSuite struct {
	test.BehaviorTestSuite
	repo *repository.NUrlRepository
	nurl *model.NUrl
}

func (suite *NurlRepositoryTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.repo = repository.NewNurlRepository(suite.AeroSpikeConnection())
	ssp := data_generator.GenerateSSP()
	dsp := data_generator.GenerateDSP()
	suite.nurl = data_generator.GenerateNUrl(ssp, dsp)
}

func (suite *NurlRepositoryTestSuite) TestCreate() {
	err := suite.repo.Create(suite.nurl)
	helper.PanicOnError(err)
}

func (suite *NurlRepositoryTestSuite) TestGetByToken() {
	suite.repo.Create(suite.nurl)

	nurlCreated, err := suite.repo.FindById(suite.nurl.Token)
	helper.PanicOnError(err)

	suite.Equal(suite.nurl.Token, nurlCreated.Token)
	suite.Equal(suite.nurl.DSPNurl, nurlCreated.DSPNurl)
	suite.Equal(suite.nurl.IsUsed, nurlCreated.IsUsed)
}

func (suite *NurlRepositoryTestSuite) TestUpdate() {
	suite.repo.Create(suite.nurl)
	suite.repo.MarkAsUsed(suite.nurl)

	nurlUpdated, err := suite.repo.FindById(suite.nurl.Token)
	helper.PanicOnError(err)

	suite.Equal(suite.nurl.Token, nurlUpdated.Token)
	suite.Equal(suite.nurl.DSPNurl, nurlUpdated.DSPNurl)
	suite.Equal(true, nurlUpdated.IsUsed)
}

func (suite *NurlRepositoryTestSuite) TearDownTest() {
	if suite.nurl != nil {
		suite.repo.Delete(suite.nurl.Token)
	}
	suite.BehaviorTestSuite.TearDownTest()
}

func TestNurlRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(NurlRepositoryTestSuite))
}
