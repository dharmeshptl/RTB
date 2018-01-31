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

type StatServiceTestSuite struct {
	test.BehaviorTestSuite
	statRepo    *repository.StatRepository
	statService *service.StatService
	stat        *model.Stat
}

func (suite *StatServiceTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.statRepo = repository.NewStatRepository(suite.AeroSpikeConnection())
	suite.statService = service.NewStatService(suite.statRepo)
}

func (suite *StatServiceTestSuite) TestLogWinConfirm() {
	ssp := data_generator.GenerateSSP()
	dsp := data_generator.GenerateDSP()
	logId := helper.GetLogId(ssp.ID, dsp.ID)
	nurl := data_generator.GenerateNUrl(
		ssp,
		dsp,
	)
	err := suite.statService.LogWinConfirm(nurl, ssp, dsp)
	helper.PanicOnError(err)

	suite.stat, err = suite.statRepo.FindByID(logId)
	helper.PanicOnError(err)
	suite.Equal(suite.stat.SSPId, ssp.ID)
	suite.Equal(suite.stat.DSPId, dsp.ID)
	suite.Equal(helper.GetCurrentStatHour(), suite.stat.StatHour)
	suite.Equal(uint64(1), suite.stat.Impression)
	suite.Equal(helper.CalculatePrice(nurl.BidPrice, nurl.ProfitMargin), suite.stat.SpendByHour)
	suite.Equal(nurl.BidPrice, suite.stat.EarnByHour)
}

func (suite *StatServiceTestSuite) TearDownTest() {
	if suite.stat != nil {
		suite.statRepo.Delete(suite.stat)
	}
	suite.BehaviorTestSuite.TearDownTest()
}

func TestStatLogsServiceRunner(t *testing.T) {
	suite.Run(t, new(StatServiceTestSuite))
}
