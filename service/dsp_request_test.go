package service_test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"go_rtb/internal/env"
	"go_rtb/internal/model"
	"go_rtb/internal/repository"
	"go_rtb/internal/service"
	"go_rtb/internal/test"
	"go_rtb/internal/test/data_generator"
	"go_rtb/internal/tool/helper"
	"go_rtb/internal/tool/validation"
	"testing"
)

type DSPRequestServiceTestSuite struct {
	test.BehaviorTestSuite
	sspRepo *repository.SSPRepository
	dspRepo *repository.DSPRepository
	service *service.DSPRequestService

	ssp     *model.SSP
	dspList []*model.DSP
}

func (suite *DSPRequestServiceTestSuite) SetupTest() {
	suite.BehaviorTestSuite.SetupTest()
	suite.sspRepo = repository.NewSSPRepository(suite.AeroSpikeConnection())
	suite.dspRepo = repository.NewDSPRepository(suite.AeroSpikeConnection())
	suite.service = service.NewDSPRequestService(
		suite.dspRepo,
		data_generator.GenerateCheckerList(),
	)

	requestEnv := env.NewEnv(context.Background(), validation.NewValidator())
	suite.service.SetEnv(requestEnv)
}

func (suite *DSPRequestServiceTestSuite) TestGetDSPToCall() {
	bidRequest := data_generator.GenerateSimpleBannerBidRequest()
	suite.ssp = data_generator.GenerateSSP()

	dspWithSupportAllConfig := data_generator.GenerateDSP()
	suite.dspList = append(suite.dspList, dspWithSupportAllConfig)
	suite.ssp.DSPIds = []uint{dspWithSupportAllConfig.ID}

	err := suite.sspRepo.Create(suite.ssp)
	helper.PanicOnError(err)

	err = suite.dspRepo.Create(dspWithSupportAllConfig)
	helper.PanicOnError(err)

	dspList, errList := suite.service.GetDSPToCall(suite.ssp, bidRequest)

	suite.Equal(0, len(errList))
	suite.Equal(1, len(dspList))
}

func (suite *DSPRequestServiceTestSuite) TearDownTest() {
	if suite.ssp != nil {
		suite.sspRepo.Delete(suite.ssp)
	}
	for _, dsp := range suite.dspList {
		suite.dspRepo.Delete(dsp)
	}

	suite.BehaviorTestSuite.TearDownTest()
}

func TestDSPRequestServiceRunner(t *testing.T) {
	suite.Run(t, new(DSPRequestServiceTestSuite))
}
