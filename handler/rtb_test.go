package handler_test

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
	"go_rtb/internal/model"
	"go_rtb/internal/protocol_buffer"
	"go_rtb/internal/repository"
	"go_rtb/internal/test"
	"go_rtb/internal/test/data_generator"
	"go_rtb/internal/tool/helper"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type RTBHandlerTestSuite struct {
	test.BehaviorTestSuite
	sspRepo *repository.SSPRepository
	dspRepo *repository.DSPRepository
	ssp     *model.SSP
	dspList []*model.DSP
}

func (suite *RTBHandlerTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.sspRepo = repository.NewSSPRepository(suite.AeroSpikeConnection())
	suite.dspRepo = repository.NewDSPRepository(suite.AeroSpikeConnection())
}

func (suite *RTBHandlerTestSuite) TestSSPNurlRequestHandler() {
	token := uuid.NewV4()
	res := suite.ApiTest().Get(
		fmt.Sprintf("/win_confirm/%s", token.String()),
	)
	suite.Equal(http.StatusNoContent, res.Result().StatusCode)
}

func (suite *RTBHandlerTestSuite) TestRtbRequestHandlerSuccess() {
	marshaler := jsonpb.Marshaler{}

	simpleBidResonse := data_generator.GenerateSimpleBidResponse()
	data, err := marshaler.MarshalToString(simpleBidResonse)
	helper.PanicOnError(err)
	body := string(data)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, body)
	}))
	defer ts.Close()

	suite.ssp = data_generator.GenerateSSP()
	suite.ssp.ProfitMargin = 7
	dspWithSupportAllConfig := data_generator.GenerateDSP()
	dspWithSupportAllConfig.EndpointURL = ts.URL
	suite.dspList = append(suite.dspList, dspWithSupportAllConfig)
	suite.ssp.DSPIds = []uint{dspWithSupportAllConfig.ID}

	err = suite.sspRepo.Create(suite.ssp)
	helper.PanicOnError(err)
	err = suite.dspRepo.Create(dspWithSupportAllConfig)
	helper.PanicOnError(err)

	bidRequest := data_generator.GenerateSimpleBannerBidRequest()
	dataBidRequest, err := marshaler.MarshalToString(bidRequest)
	helper.PanicOnError(err)
	bodyBidRequest := string(dataBidRequest)
	res := suite.ApiTest().Post(
		fmt.Sprintf("/rtb-request/%s", suite.ssp.ApiKey),
		strings.NewReader(bodyBidRequest),
	)
	var result protocol_buffer.BidResponse
	err = test.ParseResponse(res.Body, &result)
	helper.PanicOnError(err)

	suite.Equal(http.StatusOK, res.Result().StatusCode)
	suite.Equal(*simpleBidResonse.Id, *result.Id)
	suite.Equal(*simpleBidResonse.Bidid, *result.Bidid)
	suite.Equal(*simpleBidResonse.Seatbid[0].Seat, *result.Seatbid[0].Seat)
	suite.Equal(*simpleBidResonse.Seatbid[0].Bid[0].Id, *result.Seatbid[0].Bid[0].Id)
	suite.Equal(*simpleBidResonse.Seatbid[0].Bid[0].Impid, *result.Seatbid[0].Bid[0].Impid)
	suite.Equal(8.7699, *result.Seatbid[0].Bid[0].Price)
	suite.True(
		strings.Contains(
			*result.Seatbid[0].Bid[0].Nurl,
			fmt.Sprintf("%s/win_confirm/", suite.GetSystemConfig().App.BaseUrl),
		),
	)
	suite.Equal(*simpleBidResonse.Seatbid[0].Bid[0].Iurl, *result.Seatbid[0].Bid[0].Iurl)
	suite.Equal(*simpleBidResonse.Seatbid[0].Bid[0].Cid, *result.Seatbid[0].Bid[0].Cid)
	suite.Equal(*simpleBidResonse.Seatbid[0].Bid[0].Crid, *result.Seatbid[0].Bid[0].Crid)
}

func (suite *RTBHandlerTestSuite) TestRtbRequestHandlerWithNoValidDSP() {
	marshaler := jsonpb.Marshaler{}

	simpleBidResonse := data_generator.GenerateSimpleBidResponse()
	data, err := marshaler.MarshalToString(simpleBidResonse)
	helper.PanicOnError(err)
	body := string(data)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, body)
	}))
	defer ts.Close()

	suite.ssp = data_generator.GenerateSSP()
	suite.ssp.ProfitMargin = 7
	dspWithNoBannerType := data_generator.GenerateDSP()
	dspWithNoBannerType.BannerTypes = []string{}
	dspWithNoBannerType.EndpointURL = ts.URL
	suite.dspList = append(suite.dspList, dspWithNoBannerType)
	suite.ssp.DSPIds = []uint{dspWithNoBannerType.ID}

	err = suite.sspRepo.Create(suite.ssp)
	helper.PanicOnError(err)
	err = suite.dspRepo.Create(dspWithNoBannerType)
	helper.PanicOnError(err)

	bidRequest := data_generator.GenerateSimpleBannerBidRequest()
	dataBidRequest, err := marshaler.MarshalToString(bidRequest)
	helper.PanicOnError(err)
	bodyBidRequest := string(dataBidRequest)
	res := suite.ApiTest().Post(
		fmt.Sprintf("/rtb-request/%s", suite.ssp.ApiKey),
		strings.NewReader(bodyBidRequest),
	)
	var result protocol_buffer.BidResponse
	err = test.ParseResponse(res.Body, &result)
	helper.PanicOnError(err)

	suite.Equal(http.StatusOK, res.Result().StatusCode)
}

func (suite *RTBHandlerTestSuite) TearDownTest() {
	if suite.ssp != nil {
		suite.sspRepo.Delete(suite.ssp)
	}
	for _, dsp := range suite.dspList {
		suite.dspRepo.Delete(dsp)
	}

	suite.BehaviorTestSuite.TearDownTest()
}

func TestRTBHandlerRunner(t *testing.T) {
	suite.Run(t, new(RTBHandlerTestSuite))
}
