package main

import (
	"github.com/aerospike/aerospike-client-go"
	"go_rtb/internal/configuration"
	"go_rtb/internal/connection"
)

func main() {
	systemConfig, err := configuration.LoadConfig("config/setting.json")
	dbConfig := &systemConfig.DB
	if err != nil {
		panic(err)
	}

	db := connection.NewAerospikeConnection(dbConfig)

	db.AsClient().CreateIndex(
		nil,
		dbConfig.Namespace,
		dbConfig.SSPLogSet,
		"idx_gortb_ssplog_sspid",
		"ssp_id",
		aerospike.NUMERIC,
	)

	db.AsClient().CreateIndex(
		nil,
		dbConfig.Namespace,
		dbConfig.SSPLogSet,
		"idx_gortb_ssplog_hour",
		"stat_hour",
		aerospike.NUMERIC,
	)

	db.AsClient().CreateIndex(
		nil,
		dbConfig.Namespace,
		dbConfig.DSPLogSet,
		"idx_gortb_dsplog_dspid",
		"dsp_id",
		aerospike.NUMERIC,
	)

	db.AsClient().CreateIndex(
		nil,
		dbConfig.Namespace,
		dbConfig.DSPLogSet,
		"idx_gortb_dsplog_hour",
		"stat_hour",
		aerospike.NUMERIC,
	)

	db.AsClient().CreateIndex(
		nil,
		dbConfig.Namespace,
		dbConfig.StatSet,
		"idx_gortb_stat_dspid",
		"dsp_id",
		aerospike.NUMERIC,
	)

	db.AsClient().CreateIndex(
		nil,
		dbConfig.Namespace,
		dbConfig.StatSet,
		"idx_gortb_stat_sspid",
		"ssp_id",
		aerospike.NUMERIC,
	)

	db.AsClient().CreateIndex(
		nil,
		dbConfig.Namespace,
		dbConfig.StatSet,
		"idx_gortb_stat_hour",
		"start_hour",
		aerospike.NUMERIC,
	)
}
