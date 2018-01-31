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

type DSPLogsServiceTestSuite struct {
	test.BehaviorTestSuite
	dspLogRepo    *repository.DSPLogsRepository
	dspLog        *model.DSPLog
	dspLogService *service.DSPLogService
	dsp           *model.DSP
}

func (suite *DSPLogsServiceTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.dspLogRepo = repository.NewDSPLogsRepository(suite.AeroSpikeConnection())
	suite.dspLogService = service.NewDSPLogService(suite.dspLogRepo)
	suite.dsp = data_generator.GenerateDSP()
}

func (suite *DSPLogsServiceTestSuite) TestLogPreRequest() {
	logId := helper.GetLogId(suite.dsp.ID)
	err := suite.dspLogService.LogPreRequest(suite.dsp)
	helper.PanicOnError(err)

	suite.dspLog, err = suite.dspLogRepo.FindByID(logId)
	helper.PanicOnError(err)
	suite.Equal(suite.dspLog.DSPId, suite.dsp.ID)
	suite.Equal(uint64(1), suite.dspLog.SumRequest)
	suite.Equal(uint64(0), suite.dspLog.SumResponse)
	suite.Equal(float64(0), suite.dspLog.SpendByHour)
	suite.Equal(float64(0), suite.dspLog.BidQps)
	suite.NotEqual(float64(0), suite.dspLog.Qps)
}

func (suite *DSPLogsServiceTestSuite) TestLogAfterRequest() {
	logId := helper.GetLogId(suite.dsp.ID)
	err := suite.dspLogService.LogAfterRequest(suite.dsp)
	helper.PanicOnError(err)

	suite.dspLog, err = suite.dspLogRepo.FindByID(logId)
	helper.PanicOnError(err)
	suite.Equal(suite.dspLog.DSPId, suite.dsp.ID)
	suite.Equal(uint64(0), suite.dspLog.SumRequest)
	suite.Equal(uint64(1), suite.dspLog.SumResponse)
	suite.Equal(float64(0), suite.dspLog.SpendByHour)
	suite.Equal(float64(0), suite.dspLog.Qps)
	suite.NotEqual(float64(0), suite.dspLog.BidQps)
}

func (suite *DSPLogsServiceTestSuite) TestLogWinConfirm() {
	logId := helper.GetLogId(suite.dsp.ID)
	nurl := data_generator.GenerateNUrl(
		data_generator.GenerateSSP(),
		data_generator.GenerateDSP(),
	)
	err := suite.dspLogService.LogWinConfirm(nurl, suite.dsp)
	helper.PanicOnError(err)

	suite.dspLog, err = suite.dspLogRepo.FindByID(logId)
	helper.PanicOnError(err)
	suite.Equal(suite.dspLog.DSPId, suite.dsp.ID)
	suite.Equal(uint64(0), suite.dspLog.SumRequest)
	suite.Equal(uint64(0), suite.dspLog.SumResponse)
	suite.Equal(nurl.BidPrice, suite.dspLog.SpendByHour)
	suite.Equal(float64(0), suite.dspLog.Qps)
	suite.Equal(float64(0), suite.dspLog.BidQps)
}

func (suite *DSPLogsServiceTestSuite) TearDownTest() {
	if suite.dspLog != nil {
		suite.dspLogRepo.Delete(suite.dspLog)
	}
	suite.BehaviorTestSuite.TearDownTest()
}

func TestDSPLogsServiceRunner(t *testing.T) {
	suite.Run(t, new(DSPLogsServiceTestSuite))
}
