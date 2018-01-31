package repository_test

import (
	"github.com/stretchr/testify/suite"
	"go_rtb/internal/model"
	"go_rtb/internal/repository"
	"go_rtb/internal/test"
	"go_rtb/internal/test/data_generator"
	"go_rtb/internal/tool/helper"
	"go_rtb/internal/value_object"
	"testing"
)

type StatRepoTestSuite struct {
	test.BehaviorTestSuite
	statRepo *repository.StatRepository
	dspRepo  *repository.DSPRepository
	sspRepo  *repository.SSPRepository
	dsp      *model.DSP
	ssp      *model.SSP
	stat     *model.Stat
}

func (suite *StatRepoTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.statRepo = repository.NewStatRepository(suite.AeroSpikeConnection())
	suite.dspRepo = repository.NewDSPRepository(suite.AeroSpikeConnection())
	suite.sspRepo = repository.NewSSPRepository(suite.AeroSpikeConnection())
	suite.dsp = data_generator.GenerateDSP()
	suite.ssp = data_generator.GenerateSSP()
}

func (suite *StatRepoTestSuite) TestCreateForSSPAndDSP() {
	err := suite.statRepo.CreateForSSPAndDSP(suite.ssp, suite.dsp)
	helper.PanicOnError(err)
	newStatId := helper.GetLogId(suite.ssp.ID, suite.dsp.ID)

	suite.stat, err = suite.statRepo.FindByID(newStatId)
	helper.PanicOnError(err)

	suite.Equal(suite.ssp.ID, suite.stat.SSPId)
	suite.Equal(suite.dsp.ID, suite.stat.DSPId)
	suite.Equal(0, int(suite.stat.Impression))
	suite.Equal(0, int(suite.stat.SpendByHour))
	suite.Equal(0, int(suite.stat.EarnByHour))
}

func (suite *StatRepoTestSuite) TestIncreaseCounter() {
	suite.statRepo.CreateForSSPAndDSP(suite.ssp, suite.dsp)
	err := suite.statRepo.IncreaseCounter(
		suite.ssp,
		suite.dsp,
		model.StatImpressionField,
	)
	helper.PanicOnError(err)

	newStatId := helper.GetLogId(suite.ssp.ID, suite.dsp.ID)
	suite.stat, err = suite.statRepo.FindByID(newStatId)

	suite.Equal(suite.ssp.ID, suite.stat.SSPId)
	suite.Equal(suite.dsp.ID, suite.stat.DSPId)
	suite.Equal(1, int(suite.stat.Impression))
	suite.Equal(0, int(suite.stat.SpendByHour))
	suite.Equal(0, int(suite.stat.EarnByHour))
}

func (suite *StatRepoTestSuite) TestIncreaseCounterNotExistStat() {
	err := suite.statRepo.IncreaseCounter(
		suite.ssp,
		suite.dsp,
		model.StatImpressionField,
	)
	helper.PanicOnError(err)

	err = suite.statRepo.IncreaseCounter(
		suite.ssp,
		suite.dsp,
		model.StatImpressionField,
	)
	helper.PanicOnError(err)

	newStatId := helper.GetLogId(suite.ssp.ID, suite.dsp.ID)
	suite.stat, err = suite.statRepo.FindByID(newStatId)
	helper.PanicOnError(err)

	suite.Equal(suite.ssp.ID, suite.stat.SSPId)
	suite.Equal(suite.dsp.ID, suite.stat.DSPId)
	suite.Equal(2, int(suite.stat.Impression))
	suite.Equal(0, int(suite.stat.SpendByHour))
	suite.Equal(0, int(suite.stat.EarnByHour))
}

func (suite *StatRepoTestSuite) TestIncreaseMoney() {
	err := suite.statRepo.IncreaseMoney(
		suite.ssp,
		suite.dsp,
		value_object.NewMoneyPlus(model.StatEarnByHourField, 0.00023),
		value_object.NewMoneyPlus(model.StatSpendByHourField, 0.000001),
	)
	helper.PanicOnError(err)

	err = suite.statRepo.IncreaseMoney(
		suite.ssp,
		suite.dsp,
		value_object.NewMoneyPlus(model.StatSpendByHourField, 0.000001),
	)
	helper.PanicOnError(err)

	newStatId := helper.GetLogId(suite.ssp.ID, suite.dsp.ID)
	suite.stat, err = suite.statRepo.FindByID(newStatId)
	helper.PanicOnError(err)

	suite.Equal(0.00023, suite.stat.EarnByHour)
	suite.Equal(0.000002, suite.stat.SpendByHour)
}

func (suite *StatRepoTestSuite) TearDownTest() {
	if suite.dsp != nil {
		suite.dspRepo.Delete(suite.dsp)
	}
	if suite.ssp != nil {
		suite.sspRepo.Delete(suite.ssp)
	}
	if suite.stat != nil {
		suite.statRepo.Delete(suite.stat)
	}
	suite.BehaviorTestSuite.TearDownTest()
}

func TestStatRepositoryRunner(t *testing.T) {
	suite.Run(t, new(StatRepoTestSuite))
}
