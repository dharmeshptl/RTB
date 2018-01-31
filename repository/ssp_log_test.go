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

type SSPLogSRepositoryTestSuite struct {
	test.BehaviorTestSuite
	sspLogRepo *repository.SSPLogsRepository
	sspRepo    *repository.SSPRepository
	ssp        *model.SSP
	sspLog     *model.SSPLog
}

func (suite *SSPLogSRepositoryTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.sspLogRepo = repository.NewSSPLogsRepository(suite.AeroSpikeConnection())
	suite.sspRepo = repository.NewSSPRepository(suite.AeroSpikeConnection())
	suite.ssp = data_generator.GenerateSSP()
}

func (suite *SSPLogSRepositoryTestSuite) TestCreateForSSP() {
	err := suite.sspLogRepo.CreateForSSP(suite.ssp)
	helper.PanicOnError(err)
	newSSPLogId := helper.GetLogId(suite.ssp.ID)

	suite.sspLog, err = suite.sspLogRepo.FindByID(newSSPLogId)
	helper.PanicOnError(err)

	suite.Equal(suite.ssp.ID, suite.sspLog.SSPId)
	suite.Equal(0, int(suite.sspLog.BidQps))
	suite.Equal(0, int(suite.sspLog.Qps))
	suite.Equal(0, int(suite.sspLog.Impression))
	suite.Equal(0, int(suite.sspLog.SumRequest))
	suite.Equal(0, int(suite.sspLog.SumResponse))
	suite.Equal(0, int(suite.sspLog.SpendByHour))
}

func (suite *SSPLogSRepositoryTestSuite) TestIncreaseCounter() {
	suite.sspLogRepo.CreateForSSP(suite.ssp)
	err := suite.sspLogRepo.IncreaseCounter(
		suite.ssp,
		model.SSPLogImpressionField,
		model.SSPLogSumRequestField,
		model.SSPLogSumResponseField,
	)
	helper.PanicOnError(err)

	newSSPLogId := helper.GetLogId(suite.ssp.ID)
	suite.sspLog, err = suite.sspLogRepo.FindByID(newSSPLogId)
	helper.PanicOnError(err)

	suite.Equal(suite.ssp.ID, suite.sspLog.SSPId)
	suite.Equal(1, int(suite.sspLog.Impression))
	suite.Equal(1, int(suite.sspLog.SumRequest))
	suite.Equal(1, int(suite.sspLog.SumResponse))
}

func (suite *SSPLogSRepositoryTestSuite) TestIncreaseCounterNotExistLog() {
	err := suite.sspLogRepo.IncreaseCounter(
		suite.ssp,
		model.SSPLogImpressionField,
		model.SSPLogSumRequestField,
		model.SSPLogSumResponseField,
	)
	helper.PanicOnError(err)

	err = suite.sspLogRepo.IncreaseCounter(
		suite.ssp,
		model.SSPLogSumResponseField,
	)
	helper.PanicOnError(err)

	newSSPLogId := helper.GetLogId(suite.ssp.ID)
	suite.sspLog, err = suite.sspLogRepo.FindByID(newSSPLogId)

	suite.Equal(suite.ssp.ID, suite.sspLog.SSPId)
	suite.Equal(1, int(suite.sspLog.Impression))
	suite.Equal(1, int(suite.sspLog.SumRequest))
	suite.Equal(2, int(suite.sspLog.SumResponse))
}

func (suite *SSPLogSRepositoryTestSuite) TestIncreaseMoney() {
	err := suite.sspLogRepo.IncreaseMoney(
		suite.ssp,
		value_object.NewMoneyPlus(model.SSPLogSpendByHourField, 0.000001),
	)
	helper.PanicOnError(err)

	err = suite.sspLogRepo.IncreaseMoney(
		suite.ssp,
		value_object.NewMoneyPlus(model.SSPLogSpendByHourField, 0.000001),
	)
	helper.PanicOnError(err)

	newStatId := helper.GetLogId(suite.ssp.ID)
	suite.sspLog, err = suite.sspLogRepo.FindByID(newStatId)
	helper.PanicOnError(err)

	suite.Equal(0.000002, suite.sspLog.SpendByHour)
}

func (suite *SSPLogSRepositoryTestSuite) TearDownTest() {
	if suite.ssp != nil {
		suite.sspRepo.Delete(suite.ssp)
	}
	if suite.sspLog != nil {
		suite.sspLogRepo.Delete(suite.sspLog)
	}
	suite.BehaviorTestSuite.TearDownTest()
}

func TestSSPLogsRepositoryRunner(t *testing.T) {
	suite.Run(t, new(SSPLogSRepositoryTestSuite))
}
