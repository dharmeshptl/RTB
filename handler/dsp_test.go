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

type DSPHandlerTestSuite struct {
	test.BehaviorTestSuite
	dspRepo *repository.DSPRepository
	dsp     *model.DSP
}

func (suite *DSPHandlerTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.dspRepo = repository.NewDSPRepository(suite.AeroSpikeConnection())
	suite.dsp = data_generator.GenerateDSP()
}

func (suite *DSPHandlerTestSuite) TestGet() {
	suite.dspRepo.Create(suite.dsp)
	res := suite.ApiTest().Get(
		fmt.Sprintf("/admin/dsp/%d", suite.dsp.ID),
	)

	var result model.DSP
	err := test.ParseResponse(res.Body, &result)
	helper.PanicOnError(err)

	suite.Equal(http.StatusOK, res.Result().StatusCode)
	suite.Equal(suite.dsp.ID, result.ID)
	suite.Equal(suite.dsp.Name, result.Name)
	suite.Equal(suite.dsp.QPSLimit, result.QPSLimit)
	suite.Equal(suite.dsp.MinFloor, result.MinFloor)
	suite.Equal(suite.dsp.Spend, result.Spend)
	suite.Equal(suite.dsp.EndpointURL, result.EndpointURL)
	suite.Equal(suite.dsp.Company, result.Company)
	suite.Equal(suite.dsp.Region, result.Region)
	suite.Equal(suite.dsp.BannerTypes, result.BannerTypes)
	suite.Equal(suite.dsp.CountryCodes, result.CountryCodes)
	suite.Equal(suite.dsp.DeviceTypes, result.DeviceTypes)
	suite.Equal(suite.dsp.AdsTypes, result.AdsTypes)
	suite.Equal(suite.dsp.AdsSizes, result.AdsSizes)
}

func (suite *DSPHandlerTestSuite) TestCreate() {
	data, err := json.Marshal(*suite.dsp)
	helper.PanicOnError(err)
	body := string(data)

	res := suite.ApiTest().Post("/admin/dsp", strings.NewReader(body))
	suite.Equal(http.StatusNoContent, res.Result().StatusCode)
	suite.False(helper.IsRecordNotFoundError(err))

	result, err := suite.dspRepo.FindByID(suite.dsp.ID)
	suite.Equal(suite.dsp.ID, result.ID)
	suite.Equal(suite.dsp.Name, result.Name)
	suite.Equal(suite.dsp.QPSLimit, result.QPSLimit)
	suite.Equal(suite.dsp.MinFloor, result.MinFloor)
	suite.Equal(suite.dsp.Spend, result.Spend)
	suite.Equal(suite.dsp.EndpointURL, result.EndpointURL)
	suite.Equal(suite.dsp.Company, result.Company)
	suite.Equal(suite.dsp.Region, result.Region)
	suite.Equal(suite.dsp.BannerTypes, result.BannerTypes)
	suite.Equal(suite.dsp.CountryCodes, result.CountryCodes)
	suite.Equal(suite.dsp.DeviceTypes, result.DeviceTypes)
	suite.Equal(suite.dsp.AdsTypes, result.AdsTypes)
	suite.Equal(suite.dsp.AdsSizes, result.AdsSizes)
}

func (suite *DSPHandlerTestSuite) TestUpdate() {
	suite.dspRepo.Create(suite.dsp)
	suite.dsp.Name = "New name"
	suite.dsp.Company = "New company"
	data, err := json.Marshal(*suite.dsp)
	helper.PanicOnError(err)
	body := string(data)

	res := suite.ApiTest().Put(
		fmt.Sprintf("/admin/dsp/%d", suite.dsp.ID),
		strings.NewReader(body),
	)
	suite.Equal(http.StatusNoContent, res.Result().StatusCode)
	suite.False(helper.IsRecordNotFoundError(err))

	result, err := suite.dspRepo.FindByID(suite.dsp.ID)
	suite.Equal(suite.dsp.ID, result.ID)
	suite.Equal(suite.dsp.Name, result.Name)
	suite.Equal(suite.dsp.QPSLimit, result.QPSLimit)
	suite.Equal(suite.dsp.MinFloor, result.MinFloor)
	suite.Equal(suite.dsp.Spend, result.Spend)
	suite.Equal(suite.dsp.EndpointURL, result.EndpointURL)
	suite.Equal(suite.dsp.Company, result.Company)
	suite.Equal(suite.dsp.Region, result.Region)
	suite.Equal(suite.dsp.BannerTypes, result.BannerTypes)
	suite.Equal(suite.dsp.CountryCodes, result.CountryCodes)
	suite.Equal(suite.dsp.DeviceTypes, result.DeviceTypes)
	suite.Equal(suite.dsp.AdsTypes, result.AdsTypes)
	suite.Equal(suite.dsp.AdsSizes, result.AdsSizes)
}

func (suite *DSPHandlerTestSuite) TestDelete() {
	suite.dspRepo.Create(suite.dsp)
	res := suite.ApiTest().Delete(
		fmt.Sprintf("/admin/dsp/%d", suite.dsp.ID),
	)

	suite.Equal(http.StatusNoContent, res.Result().StatusCode)
}

func (suite *DSPHandlerTestSuite) TearDownTest() {
	if suite.dsp != nil {
		suite.dspRepo.Delete(suite.dsp)
	}
	suite.BehaviorTestSuite.TearDownTest()
}

func TestDSPHandlerRunner(t *testing.T) {
	suite.Run(t, new(DSPHandlerTestSuite))
}
