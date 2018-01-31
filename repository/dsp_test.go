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

type DSPRepositoryTestSuite struct {
	test.BehaviorTestSuite
	repo *repository.DSPRepository
	dsp  *model.DSP
}

func (suite *DSPRepositoryTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.repo = repository.NewDSPRepository(suite.AeroSpikeConnection())
	suite.dsp = data_generator.GenerateDSP()
}

func (suite *DSPRepositoryTestSuite) TestFindById() {
	err := suite.repo.Create(suite.dsp)
	helper.PanicOnError(err)

	dspCreated, err := suite.repo.FindByID(suite.dsp.ID)
	helper.PanicOnError(err)

	suite.Equal(suite.dsp.ID, dspCreated.ID)
	suite.Equal(suite.dsp.Name, dspCreated.Name)
	suite.Equal(suite.dsp.QPSLimit, dspCreated.QPSLimit)
	suite.Equal(suite.dsp.MinFloor, dspCreated.MinFloor)
	suite.Equal(suite.dsp.Spend, dspCreated.Spend)
	suite.Equal(suite.dsp.EndpointURL, dspCreated.EndpointURL)
	suite.Equal(suite.dsp.Company, dspCreated.Company)
	suite.Equal(suite.dsp.Region, dspCreated.Region)
	suite.Equal(suite.dsp.BannerTypes, dspCreated.BannerTypes)
	suite.Equal(suite.dsp.CountryCodes, dspCreated.CountryCodes)
	suite.Equal(suite.dsp.DeviceTypes, dspCreated.DeviceTypes)
	suite.Equal(suite.dsp.AdsTypes, dspCreated.AdsTypes)
	suite.Equal(suite.dsp.AdsSizes, dspCreated.AdsSizes)
}

func (suite *DSPRepositoryTestSuite) TestDelete() {
	suite.repo.Create(suite.dsp)
	suite.repo.Delete(suite.dsp)

	_, err := suite.repo.FindByID(suite.dsp.ID)

	suite.True(helper.IsRecordNotFoundError(err))
}

func (suite *DSPRepositoryTestSuite) TestUpdate() {
	suite.repo.Create(suite.dsp)
	suite.dsp.Name = "UPdated_name"
	suite.repo.Update(suite.dsp)

	updatedDsp, _ := suite.repo.FindByID(suite.dsp.ID)

	suite.Equal(suite.dsp.ID, updatedDsp.ID)
	suite.Equal(suite.dsp.Name, updatedDsp.Name)
}

func (suite *DSPRepositoryTestSuite) TearDownTest() {
	if suite.dsp != nil {
		suite.repo.Delete(suite.dsp)
	}
	suite.BehaviorTestSuite.TearDownTest()
}

func TestDSPRepositoryRunner(t *testing.T) {
	suite.Run(t, new(DSPRepositoryTestSuite))
}
