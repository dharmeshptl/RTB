package main

import (
	"fmt"
	"go_rtb/internal/configuration"
	"go_rtb/internal/connection"
	"go_rtb/internal/message_queue/rabbitmq"
	"go_rtb/internal/message_queue/worker"
	"go_rtb/internal/repository"
	"go_rtb/internal/service"
	"go_rtb/internal/tool/logger"
)

func main() {
	systemConfig, err := configuration.LoadConfig("config/setting.json")
	if err != nil {
		panic(err)
	}
	conn := &rabbitmq.Connector{}
	err = conn.Connect(systemConfig.MessageQueue.ConnectionString)
	if err != nil {
		panic(err)
	}

	numberOfCallbackWorker := systemConfig.MessageQueue.WinConfirmConsumerNumber

	db := connection.NewAerospikeConnection(&systemConfig.DB)
	sspRepo := repository.NewSSPRepository(db)
	dspRepo := repository.NewDSPRepository(db)
	nurlRepo := repository.NewNurlRepository(db)
	sspLogRepo := repository.NewSSPLogsRepository(db)
	dspLogRepo := repository.NewDSPLogsRepository(db)
	statRepo := repository.NewStatRepository(db)
	sspLogService := service.NewSSPLogService(sspLogRepo)
	dspLogService := service.NewDSPLogService(dspLogRepo)
	statService := service.NewStatService(statRepo)

	logger.Info(
		fmt.Sprintf("Start %d consumers for winconfirm queue", numberOfCallbackWorker),
	)

	forever := make(chan bool)
	for i := 0; i < numberOfCallbackWorker; i++ {
		go func() {
			callbackWorker := worker.NewWinConfirmWorker(
				nurlRepo,
				statService,
				sspLogService,
				dspLogService,
				sspRepo,
				dspRepo,
			)
			consum := &rabbitmq.Consumer{}
			consum.Init(conn)
			consum.Consume(systemConfig.MessageQueue.WinConfirmQueue, callbackWorker)
		}()
	}
	<-forever
}
