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

type DSPLogsRepositoryTestSuite struct {
	test.BehaviorTestSuite
	dspLogRepo *repository.DSPLogsRepository
	dspRepo    *repository.DSPRepository
	dsp        *model.DSP
	dspLog     *model.DSPLog
}

func (suite *DSPLogsRepositoryTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.dspLogRepo = repository.NewDSPLogsRepository(suite.AeroSpikeConnection())
	suite.dspRepo = repository.NewDSPRepository(suite.AeroSpikeConnection())
	suite.dsp = data_generator.GenerateDSP()
}

func (suite *DSPLogsRepositoryTestSuite) TestCreateForDSP() {
	err := suite.dspLogRepo.CreateForDSP(suite.dsp)
	helper.PanicOnError(err)
	newDSPLogId := helper.GetLogId(suite.dsp.ID)

	suite.dspLog, err = suite.dspLogRepo.FindByID(newDSPLogId)
	helper.PanicOnError(err)

	suite.Equal(suite.dsp.ID, suite.dspLog.DSPId)
	suite.Equal(0, int(suite.dspLog.BidQps))
	suite.Equal(0, int(suite.dspLog.Qps))
	suite.Equal(0, int(suite.dspLog.SumRequest))
	suite.Equal(0, int(suite.dspLog.SumResponse))
	suite.Equal(0, int(suite.dspLog.SpendByHour))
}

func (suite *DSPLogsRepositoryTestSuite) TestIncreaseCounter() {
	suite.dspLogRepo.CreateForDSP(suite.dsp)
	err := suite.dspLogRepo.IncreaseCounter(
		suite.dsp,
		model.DSPLogSumRequestField,
		model.DSPLogSumResponseField,
	)
	helper.PanicOnError(err)

	newSSPLogId := helper.GetLogId(suite.dsp.ID)
	suite.dspLog, err = suite.dspLogRepo.FindByID(newSSPLogId)

	suite.Equal(suite.dsp.ID, suite.dspLog.DSPId)
	suite.Equal(1, int(suite.dspLog.SumRequest))
	suite.Equal(1, int(suite.dspLog.SumResponse))
}

func (suite *DSPLogsRepositoryTestSuite) TestIncreaseCounterNotExistLog() {
	suite.dspLogRepo.CreateForDSP(suite.dsp)
	err := suite.dspLogRepo.IncreaseCounter(
		suite.dsp,
		model.DSPLogSumRequestField,
		model.DSPLogSumResponseField,
	)
	helper.PanicOnError(err)

	err = suite.dspLogRepo.IncreaseCounter(
		suite.dsp,
		model.DSPLogSumRequestField,
	)
	helper.PanicOnError(err)

	newSSPLogId := helper.GetLogId(suite.dsp.ID)
	suite.dspLog, err = suite.dspLogRepo.FindByID(newSSPLogId)
	helper.PanicOnError(err)

	suite.Equal(suite.dsp.ID, suite.dspLog.DSPId)
	suite.Equal(2, int(suite.dspLog.SumRequest))
	suite.Equal(1, int(suite.dspLog.SumResponse))
}

func (suite *DSPLogsRepositoryTestSuite) TestIncreaseMoney() {
	err := suite.dspLogRepo.IncreaseMoney(
		suite.dsp,
		value_object.NewMoneyPlus(model.DSPLogSpendByHourField, 0.000001),
	)
	helper.PanicOnError(err)

	err = suite.dspLogRepo.IncreaseMoney(
		suite.dsp,
		value_object.NewMoneyPlus(model.DSPLogSpendByHourField, 0.000001),
	)
	helper.PanicOnError(err)

	newStatId := helper.GetLogId(suite.dsp.ID)
	suite.dspLog, err = suite.dspLogRepo.FindByID(newStatId)
	helper.PanicOnError(err)

	suite.Equal(0.000002, suite.dspLog.SpendByHour)
}

func (suite *DSPLogsRepositoryTestSuite) TearDownTest() {
	if suite.dsp != nil {
		suite.dspRepo.Delete(suite.dsp)
	}
	if suite.dspLog != nil {
		suite.dspLogRepo.Delete(suite.dspLog)
	}
	suite.BehaviorTestSuite.TearDownTest()
}

func TestDSPLogsRepositoryRunner(t *testing.T) {
	suite.Run(t, new(DSPLogsRepositoryTestSuite))
}
