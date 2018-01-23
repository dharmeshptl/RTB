package handler_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"go_rtb/internal/model"
	"go_rtb/internal/repository"
	"go_rtb/internal/test"
	"go_rtb/internal/test/data_generator"
	"go_rtb/internal/tool/helper"
	"net/http"
	"strings"
	"testing"
)

type SSPHandlerTestSuite struct {
	test.BehaviorTestSuite
	sspRepo *repository.SSPRepository
	ssp     *model.SSP
}

func (suite *SSPHandlerTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.sspRepo = repository.NewSSPRepository(suite.AeroSpikeConnection())
	suite.ssp = data_generator.GenerateSSP()
}

func (suite *SSPHandlerTestSuite) TestGet() {
	suite.sspRepo.Create(suite.ssp)
	res := suite.ApiTest().Get(
		fmt.Sprintf("/admin/ssp/%d", suite.ssp.ID),
	)

	var result model.SSP
	err := test.ParseResponse(res.Body, &result)
	helper.PanicOnError(err)

	suite.Equal(http.StatusOK, res.Result().StatusCode)
	suite.Equal(suite.ssp.ID, result.ID)
	suite.Equal(suite.ssp.Name, result.Name)
	suite.Equal(suite.ssp.ApiKey, result.ApiKey)
	suite.Equal(suite.ssp.Company, result.Company)
	suite.Equal(suite.ssp.Region, result.Region)
	suite.Equal(suite.ssp.MinFloor, result.MinFloor)
	suite.Equal(suite.ssp.ProfitMargin, result.ProfitMargin)
	suite.Equal(suite.ssp.Active, result.Active)
	suite.Equal(suite.ssp.DSPIds, result.DSPIds)
}

func (suite *SSPHandlerTestSuite) TestCreate() {
	data, err := json.Marshal(*suite.ssp)
	helper.PanicOnError(err)
	body := string(data)

	res := suite.ApiTest().Post("/admin/ssp", strings.NewReader(body))
	suite.Equal(http.StatusNoContent, res.Result().StatusCode)
	suite.False(helper.IsRecordNotFoundError(err))

	result, err := suite.sspRepo.FindByID(suite.ssp.ID)
	suite.Equal(suite.ssp.ID, result.ID)
	suite.Equal(suite.ssp.Name, result.Name)
	suite.Equal(suite.ssp.ApiKey, result.ApiKey)
	suite.Equal(suite.ssp.Company, result.Company)
	suite.Equal(suite.ssp.Region, result.Region)
	suite.Equal(suite.ssp.MinFloor, result.MinFloor)
	suite.Equal(suite.ssp.ProfitMargin, result.ProfitMargin)
	suite.Equal(suite.ssp.Active, result.Active)
	suite.Equal(suite.ssp.DSPIds, result.DSPIds)
}

func (suite *SSPHandlerTestSuite) TestUpdate() {
	suite.sspRepo.Create(suite.ssp)
	suite.ssp.Name = "New name"
	suite.ssp.Company = "New company"
	data, err := json.Marshal(*suite.ssp)
	helper.PanicOnError(err)
	body := string(data)

	res := suite.ApiTest().Put(
		fmt.Sprintf("/admin/ssp/%d", suite.ssp.ID),
		strings.NewReader(body),
	)
	suite.Equal(http.StatusNoContent, res.Result().StatusCode)
	suite.False(helper.IsRecordNotFoundError(err))

	result, err := suite.sspRepo.FindByID(suite.ssp.ID)
	suite.Equal(suite.ssp.ID, result.ID)
	suite.Equal(suite.ssp.Name, result.Name)
	suite.Equal(suite.ssp.ApiKey, result.ApiKey)
	suite.Equal(suite.ssp.Company, result.Company)
	suite.Equal(suite.ssp.Region, result.Region)
	suite.Equal(suite.ssp.MinFloor, result.MinFloor)
	suite.Equal(suite.ssp.ProfitMargin, result.ProfitMargin)
	suite.Equal(suite.ssp.Active, result.Active)
	suite.Equal(suite.ssp.DSPIds, result.DSPIds)
}

func (suite *SSPHandlerTestSuite) TestDelete() {
	suite.sspRepo.Create(suite.ssp)
	res := suite.ApiTest().Delete(
		fmt.Sprintf("/admin/ssp/%d", suite.ssp.ID),
	)

	suite.Equal(http.StatusNoContent, res.Result().StatusCode)
}

func (suite *SSPHandlerTestSuite) TearDownTest() {
	if suite.ssp != nil {
		suite.sspRepo.Delete(suite.ssp)
	}
	suite.BehaviorTestSuite.TearDownTest()
}

func TestSSPHandlerRunner(t *testing.T) {
	suite.Run(t, new(SSPHandlerTestSuite))
}
