package service_test

import (
	"github.com/stretchr/testify/suite"
	"go_rtb/internal/model"
	"go_rtb/internal/repository"
	"go_rtb/internal/service"
	"go_rtb/internal/test"
	"go_rtb/internal/test/data_generator"
	"go_rtb/internal/tool/helper"
	"testing"
)

type SSPLogsServiceTestSuite struct {
	test.BehaviorTestSuite
	sspLogRepo    *repository.SSPLogsRepository
	sspLog        *model.SSPLog
	sspLogService *service.SSPLogService
	ssp           *model.SSP
}

func (suite *SSPLogsServiceTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.sspLogRepo = repository.NewSSPLogsRepository(suite.AeroSpikeConnection())
	suite.sspLogService = service.NewSSPLogService(suite.sspLogRepo)
	suite.ssp = data_generator.GenerateSSP()
}

func (suite *SSPLogsServiceTestSuite) TestLogPreRequest() {
	logId := helper.GetLogId(suite.ssp.ID)
	err := suite.sspLogService.LogPreRequest(suite.ssp)
	helper.PanicOnError(err)

	suite.sspLog, err = suite.sspLogRepo.FindByID(logId)
	helper.PanicOnError(err)
	suite.Equal(suite.sspLog.SSPId, suite.ssp.ID)
	suite.Equal(uint64(1), suite.sspLog.SumRequest)
	suite.Equal(uint64(1), suite.sspLog.Impression)
	suite.Equal(uint64(0), suite.sspLog.SumResponse)
	suite.Equal(float64(0), suite.sspLog.SpendByHour)
	suite.Equal(float64(0), suite.sspLog.BidQps)
	suite.NotEqual(float64(0), suite.sspLog.Qps)
}

func (suite *SSPLogsServiceTestSuite) TestLogAfterRequest() {
	logId := helper.GetLogId(suite.ssp.ID)
	err := suite.sspLogService.LogAfterRequest(suite.ssp)
	helper.PanicOnError(err)

	suite.sspLog, err = suite.sspLogRepo.FindByID(logId)
	helper.PanicOnError(err)
	suite.Equal(suite.sspLog.SSPId, suite.ssp.ID)
	suite.Equal(uint64(0), suite.sspLog.SumRequest)
	suite.Equal(uint64(0), suite.sspLog.Impression)
	suite.Equal(uint64(1), suite.sspLog.SumResponse)
	suite.Equal(float64(0), suite.sspLog.SpendByHour)
	suite.Equal(float64(0), suite.sspLog.Qps)
	suite.NotEqual(float64(0), suite.sspLog.BidQps)
}

func (suite *SSPLogsServiceTestSuite) TestLogWinConfirm() {
	logId := helper.GetLogId(suite.ssp.ID)
	nurl := data_generator.GenerateNUrl(
		data_generator.GenerateSSP(),
		data_generator.GenerateDSP(),
	)
	err := suite.sspLogService.LogWinConfirm(nurl, suite.ssp)
	helper.PanicOnError(err)

	suite.sspLog, err = suite.sspLogRepo.FindByID(logId)
	helper.PanicOnError(err)
	suite.Equal(suite.sspLog.SSPId, suite.ssp.ID)
	suite.Equal(uint64(0), suite.sspLog.SumRequest)
	suite.Equal(uint64(0), suite.sspLog.SumResponse)
	suite.Equal(helper.CalculatePrice(nurl.BidPrice, nurl.ProfitMargin), suite.sspLog.SpendByHour)
	suite.Equal(float64(0), suite.sspLog.Qps)
	suite.Equal(float64(0), suite.sspLog.BidQps)
}

func (suite *SSPLogsServiceTestSuite) TearDownTest() {
	if suite.sspLog != nil {
		suite.sspLogRepo.Delete(suite.sspLog)
	}
	suite.BehaviorTestSuite.TearDownTest()
}

func TestSSPLogsServiceRunner(t *testing.T) {
	suite.Run(t, new(SSPLogsServiceTestSuite))
}
