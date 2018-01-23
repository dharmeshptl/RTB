package test

import (
	"github.com/stretchr/testify/suite"
	"go_rtb/internal/configuration"
	"go_rtb/internal/connection"
	"go_rtb/internal/router"
)

type BehaviorTestSuite struct {
	suite.Suite
	asConnection connection.AeroSpikeConnection
	systemConfig *configuration.SystemConfig
}

func (suite *BehaviorTestSuite) SetupTest() {
	var err error
	suite.systemConfig, err = configuration.LoadConfig("../../config/setting.json")
	if err != nil {
		panic(err)
	}

	suite.asConnection = connection.NewAerospikeConnection(&suite.systemConfig.DB)
}

func (suite *BehaviorTestSuite) TearDownTest() {
	suite.asConnection.AsClient().Close()
}

func (suite *BehaviorTestSuite) AeroSpikeConnection() connection.AeroSpikeConnection {
	return suite.asConnection
}

func (suite *BehaviorTestSuite) GetSystemConfig() *configuration.SystemConfig {
	return suite.systemConfig
}

func (suite *BehaviorTestSuite) ApiTest() *ApiTest {
	r := router.Init(suite.systemConfig)
	return NewApiTest(r)
}
