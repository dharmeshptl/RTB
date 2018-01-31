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

type SSPRepositoryTestSuite struct {
	test.BehaviorTestSuite
	repo *repository.SSPRepository
	ssp  *model.SSP
}

func (suite *SSPRepositoryTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.repo = repository.NewSSPRepository(suite.AeroSpikeConnection())
	suite.ssp = data_generator.GenerateSSP()
}

func (suite *SSPRepositoryTestSuite) TestFindById() {
	err := suite.repo.Create(suite.ssp)
	helper.PanicOnError(err)

	sspCreated, err := suite.repo.FindByID(suite.ssp.ID)
	helper.PanicOnError(err)

	suite.Equal(suite.ssp.ID, sspCreated.ID)
	suite.Equal(suite.ssp.Name, sspCreated.Name)
	suite.Equal(suite.ssp.ApiKey, sspCreated.ApiKey)
	suite.Equal(suite.ssp.Company, sspCreated.Company)
	suite.Equal(suite.ssp.Region, sspCreated.Region)
	suite.Equal(suite.ssp.MinFloor, sspCreated.MinFloor)
	suite.Equal(suite.ssp.ProfitMargin, sspCreated.ProfitMargin)
	suite.Equal(suite.ssp.Active, sspCreated.Active)
	suite.Equal(suite.ssp.DSPIds, sspCreated.DSPIds)
}

func (suite *SSPRepositoryTestSuite) TestDelete() {
	suite.repo.Create(suite.ssp)
	suite.repo.Delete(suite.ssp)

	_, err := suite.repo.FindByID(suite.ssp.ID)

	suite.True(helper.IsRecordNotFoundError(err))
}

func (suite *SSPRepositoryTestSuite) TestUpdate() {
	suite.repo.Create(suite.ssp)
	suite.ssp.Name = "UPdated_name"
	suite.repo.Update(suite.ssp)

	updatedSsp, _ := suite.repo.FindByID(suite.ssp.ID)

	suite.Equal(suite.ssp.ID, updatedSsp.ID)
	suite.Equal(suite.ssp.Name, updatedSsp.Name)
}

func (suite *SSPRepositoryTestSuite) TearDownTest() {
	if suite.ssp != nil {
		suite.repo.Delete(suite.ssp)
	}
	suite.BehaviorTestSuite.TearDownTest()
}

func TestSSPRepositoryRunner(t *testing.T) {
	suite.Run(t, new(SSPRepositoryTestSuite))
}
