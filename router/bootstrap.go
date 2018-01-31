package router

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"go_rtb/internal/configuration"
	"go_rtb/internal/connection"
	"go_rtb/internal/env"
	"go_rtb/internal/handler"
	"go_rtb/internal/message_queue/rabbitmq"
	"go_rtb/internal/repository"
	"go_rtb/internal/router/middleware"
	"go_rtb/internal/rtb/request_checker"
	"go_rtb/internal/service"
	"go_rtb/internal/tool/logger"
	"go_rtb/internal/tool/validation"
	"go_rtb/internal/transport"
	"go_rtb/internal/worker_pool"
	"go_rtb/internal/worker_pool/job"
)

var db connection.AeroSpikeConnection
var sspRepo *repository.SSPRepository
var dspRepo *repository.DSPRepository
var nurlRepo *repository.NUrlRepository
var sspLogRepo *repository.SSPLogsRepository
var dspLogRepo *repository.DSPLogsRepository
var statRepo *repository.StatRepository

func Init(systemConfig *configuration.SystemConfig) chi.Router {
	flag.Parse()

	db = connection.NewAerospikeConnection(&systemConfig.DB)
	sspRepo = repository.NewSSPRepository(db)
	dspRepo = repository.NewDSPRepository(db)
	nurlRepo = repository.NewNurlRepository(db)
	sspLogRepo = repository.NewSSPLogsRepository(db)
	dspLogRepo = repository.NewDSPLogsRepository(db)
	statRepo = repository.NewStatRepository(db)

	conn := &rabbitmq.Connector{}
	err := conn.Connect(systemConfig.MessageQueue.ConnectionString)
	if err != nil {
		panic(err)
	}
	messagePublisher := &rabbitmq.Publisher{}
	messagePublisher.Init(conn)

	requestCaller := &transport.HttpCaller{}
	requestCaller.Init()

	callbackService := service.NewSSPCallbackService(
		nurlRepo,
		&systemConfig.App,
	)

	checkerList := []request_checker.DSPRequestChecker{
		&request_checker.AdsSizeChecker{},
		&request_checker.AdsTypeChecker{},
		&request_checker.BannerTypeChecker{},
		&request_checker.CountryChecker{},
		&request_checker.DeviceTypeChecker{},
	}
	dspRequestService := service.NewDSPRequestService(
		dspRepo,
		checkerList,
	)

	sspLogService := service.NewSSPLogService(sspLogRepo)
	dspLogService := service.NewDSPLogService(dspLogRepo)

	rtbHandler := handler.NewRtbRequestHandler(
		requestCaller,
		systemConfig,
		callbackService,
		dspRequestService,
		sspLogService,
		dspLogService,
		messagePublisher,
	)

	startQueue(systemConfig.JobQueue)

	// Validation
	validator := validation.NewValidator()
	envFactory := func(ctx context.Context) *env.Env {
		return env.NewEnv(ctx, validator)
	}
	r := chi.NewRouter()
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.RequestID)
	r.Use(middleware.Env(envFactory))

	//Handle client rtb request
	r.With(middleware.SSPApiKeyContext(sspRepo), middleware.RtbRequestMiddleware).
		Post("/rtb-request/{sspApiKey}", middleware.MakeHandler(rtbHandler.HandleRtbRequest))

	r.Get("/win_confirm/{token}", middleware.MakeHandler(rtbHandler.HandleSSPNurl))

	r.Mount("/admin", adminRouter())

	return r
}

func adminRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.AdminProtector)

	sspHandler := handler.NewSSPHandler(sspRepo, dspRepo)
	r.Route("/ssp", func(r chi.Router) {
		r.Post("/", middleware.MakeHandler(sspHandler.Create))
		r.Route("/{sspId}", func(r chi.Router) {
			r.Use(middleware.SSPContext(sspRepo))
			r.Get("/", middleware.MakeHandler(sspHandler.Get))
			r.Put("/", middleware.MakeHandler(sspHandler.Update))
			r.Delete("/", middleware.MakeHandler(sspHandler.Delete))
		})
	})

	dspHandler := handler.NewDSPApiHandler(sspRepo, dspRepo)
	r.Route("/dsp", func(r chi.Router) {
		r.Post("/", middleware.MakeHandler(dspHandler.Create))
		r.Route("/{dspId}", func(r chi.Router) {
			r.Use(middleware.DSPContext(dspRepo))
			r.Get("/", middleware.MakeHandler(dspHandler.Get))
			r.Put("/", middleware.MakeHandler(dspHandler.Update))
			r.Delete("/", middleware.MakeHandler(dspHandler.Delete))
		})
	})

	logHandler := handler.NewLogHandler(
		sspLogRepo,
		dspLogRepo,
		statRepo,
	)

	r.Get("/ssp_log", middleware.MakeHandler(logHandler.GetSSPLog))
	r.Get("/dsp_log", middleware.MakeHandler(logHandler.GetDSPLog))
	r.Get("/stat", middleware.MakeHandler(logHandler.GetStat))

	return r
}

func startQueue(config configuration.JobQueueConfig) {
	logger.Info(fmt.Sprintf("Creating queue length %v, and %v workers", config.QueueLength, config.WorkerNumber))

	worker_pool.RtbApiCallJobQueue = make(chan *job.DSPCallJob, config.QueueLength)
	jobDispatcher := worker_pool.NewDispatcher(config.WorkerNumber)
	jobDispatcher.Run()
}
