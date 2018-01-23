package handler_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"go_rtb/internal/model"
	"go_rtb/internal/repository"
	"go_rtb/internal/test"
	"go_rtb/internal/test/data_generator"
	"go_rtb/internal/tool/helper"
	"net/http"
	"testing"
)

type LogHandlerTestSuite struct {
	test.BehaviorTestSuite
	sspLogRepo *repository.SSPLogsRepository
	dspLogRepo *repository.DSPLogsRepository
	statRepo   *repository.StatRepository
	sspLog     *model.SSPLog
	dspLog     *model.DSPLog
	stat       *model.Stat
}

func (suite *LogHandlerTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.sspLogRepo = repository.NewSSPLogsRepository(suite.AeroSpikeConnection())
	suite.dspLogRepo = repository.NewDSPLogsRepository(suite.AeroSpikeConnection())
	suite.statRepo = repository.NewStatRepository(suite.AeroSpikeConnection())
}

func (suite *LogHandlerTestSuite) TestGetSSPLog() {
	ssp := data_generator.GenerateSSP()
	err := suite.sspLogRepo.CreateForSSP(ssp)
	helper.PanicOnError(err)

	res := suite.ApiTest().Get(
		fmt.Sprintf("/admin/ssp_log?ssp_id=%d", ssp.ID),
	)

	suite.sspLog, err = suite.sspLogRepo.FindByID(helper.GetLogId(ssp.ID))
	helper.PanicOnError(err)

	var result model.SSPLog
	err = test.ParseResponse(res.Body, &result)
	helper.PanicOnError(err)
	suite.Equal(http.StatusOK, res.Result().StatusCode)
	suite.Equal(suite.sspLog.ID, result.ID)
	suite.Equal(suite.sspLog.SSPId, result.SSPId)
	suite.Equal(suite.sspLog.StatHour, result.StatHour)
	suite.Equal(suite.sspLog.BidQps, result.BidQps)
	suite.Equal(suite.sspLog.Qps, result.Qps)
	suite.Equal(suite.sspLog.Impression, result.Impression)
	suite.Equal(suite.sspLog.SumRequest, result.SumRequest)
	suite.Equal(suite.sspLog.SumResponse, result.SumResponse)
	suite.Equal(suite.sspLog.SpendByHour, result.SpendByHour)
}

func (suite *LogHandlerTestSuite) TestGetDSPLog() {
	dsp := data_generator.GenerateDSP()
	err := suite.dspLogRepo.CreateForDSP(dsp)
	helper.PanicOnError(err)

	res := suite.ApiTest().Get(
		fmt.Sprintf("/admin/dsp_log?dsp_id=%d", dsp.ID),
	)

	suite.dspLog, err = suite.dspLogRepo.FindByID(helper.GetLogId(dsp.ID))
	helper.PanicOnError(err)

	var result model.DSPLog
	err = test.ParseResponse(res.Body, &result)
	helper.PanicOnError(err)
	suite.Equal(http.StatusOK, res.Result().StatusCode)
	suite.Equal(suite.dspLog.ID, result.ID)
	suite.Equal(suite.dspLog.DSPId, result.DSPId)
	suite.Equal(suite.dspLog.StatHour, result.StatHour)
	suite.Equal(suite.dspLog.BidQps, result.BidQps)
	suite.Equal(suite.dspLog.Qps, result.Qps)
	suite.Equal(suite.dspLog.SumRequest, result.SumRequest)
	suite.Equal(suite.dspLog.SumResponse, result.SumResponse)
	suite.Equal(suite.dspLog.SpendByHour, result.SpendByHour)
}

func (suite *LogHandlerTestSuite) TestGetStat() {
	ssp := data_generator.GenerateSSP()
	dsp := data_generator.GenerateDSP()
	err := suite.statRepo.CreateForSSPAndDSP(ssp, dsp)
	helper.PanicOnError(err)

	res := suite.ApiTest().Get(
		fmt.Sprintf("/admin/stat?dsp_id=%d&ssp_id=%d", dsp.ID, ssp.ID),
	)

	suite.stat, err = suite.statRepo.FindByID(helper.GetLogId(ssp.ID, dsp.ID))
	helper.PanicOnError(err)

	var result model.Stat
	err = test.ParseResponse(res.Body, &result)
	helper.PanicOnError(err)
	suite.Equal(http.StatusOK, res.Result().StatusCode)
	suite.Equal(suite.stat.ID, result.ID)
	suite.Equal(suite.stat.SSPId, result.SSPId)
	suite.Equal(suite.stat.DSPId, result.DSPId)
	suite.Equal(suite.stat.StatHour, result.StatHour)
	suite.Equal(suite.stat.Impression, result.Impression)
	suite.Equal(suite.stat.SpendByHour, result.SpendByHour)
	suite.Equal(suite.stat.EarnByHour, result.EarnByHour)
}

func (suite *LogHandlerTestSuite) TearDownTest() {
	if suite.sspLog != nil {
		suite.sspLogRepo.Delete(suite.sspLog)
	}
	if suite.dspLog != nil {
		suite.dspLogRepo.Delete(suite.dspLog)
	}
	if suite.stat != nil {
		suite.statRepo.Delete(suite.stat)
	}
	suite.BehaviorTestSuite.TearDownTest()
}

func TestLogHandlerRunner(t *testing.T) {
	suite.Run(t, new(LogHandlerTestSuite))
}
