package service_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"go_rtb/internal/repository"
	"go_rtb/internal/service"
	"go_rtb/internal/test"
	"go_rtb/internal/test/data_generator"
	"go_rtb/internal/tool/helper"
	"strings"
	"testing"
)

type SSPCallbackServiceTestSuite struct {
	test.BehaviorTestSuite
	nurlRepo *repository.NUrlRepository
	service  *service.SSPCallbackService
}

func (suite *SSPCallbackServiceTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.nurlRepo = repository.NewNurlRepository(suite.AeroSpikeConnection())
	suite.service = service.NewSSPCallbackService(
		suite.nurlRepo,
		&suite.GetSystemConfig().App,
	)
}

func (suite *SSPCallbackServiceTestSuite) TestBuildSSPNUrl() {
	result, err := suite.service.BuildSSPNUrl(
		data_generator.GeneraSimpleBidResult(),
		data_generator.GenerateSSP(),
	)
	helper.PanicOnError(err)

	suite.True(
		strings.Contains(
			result,
			fmt.Sprintf("%s/win_confirm/", suite.GetSystemConfig().App.BaseUrl),
		),
	)
}

func (suite *SSPCallbackServiceTestSuite) TearDownTest() {
	suite.BehaviorTestSuite.TearDownTest()
}

func TestSSPCallbackServiceRunner(t *testing.T) {
	suite.Run(t, new(SSPCallbackServiceTestSuite))
}
