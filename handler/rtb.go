package handler

import (
	"bytes"
	"github.com/golang/protobuf/jsonpb"
	"go_rtb/internal/configuration"
	"go_rtb/internal/handler/response"

	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"go_rtb/internal/message_queue/payload"
	"go_rtb/internal/message_queue/rabbitmq"
	"go_rtb/internal/protocol_buffer"
	"go_rtb/internal/router/middleware"
	"go_rtb/internal/rtb"
	"go_rtb/internal/service"
	"go_rtb/internal/tool"
	"go_rtb/internal/tool/logger"
	"go_rtb/internal/transport"
	"go_rtb/internal/worker_pool"
	"go_rtb/internal/worker_pool/job"
	"io/ioutil"
	"net/http"
	"sync"
	"fmt"
)

type RtbHandler struct {
	requestCaller     *transport.HttpCaller
	systemConfig      *configuration.SystemConfig
	callbackService   *service.SSPCallbackService
	dspRequestService *service.DSPRequestService
	sspLogService     *service.SSPLogService
	dspLogService     *service.DSPLogService
	messagePublisher  *rabbitmq.Publisher
}

func NewRtbRequestHandler(
	caller *transport.HttpCaller,
	systemConfig *configuration.SystemConfig,
	callbackService *service.SSPCallbackService,
	dspService *service.DSPRequestService,
	sspLogService *service.SSPLogService,
	dspLogService *service.DSPLogService,
	messagePublisher *rabbitmq.Publisher,
) RtbHandler {
	return RtbHandler{caller, systemConfig, callbackService, dspService, sspLogService, dspLogService, messagePublisher}
}

//This is pretty bad
func (h RtbHandler) HandleRtbRequest(ctx tool.AppContext) response.ApiResponse {
	data, err := ioutil.ReadAll(ctx.GetRequest().Body)
	if err != nil {
		return response.ErrorResponse(err, http.StatusUnprocessableEntity)
	}

	ioReaderForJson := bytes.NewReader(data)
	bidResquest := protocol_buffer.BidRequest{}
	err = jsonpb.Unmarshal(ioReaderForJson, &bidResquest)
	if err != nil {
		return response.ErrorResponse(err, http.StatusUnprocessableEntity)
	}

	ssp := middleware.GetSSPFromContextApiKey(ctx.GetRequest().Context())
	h.sspLogService.LogPreRequest(ssp)

	var rtbApiCallJob job.DSPCallJob
	//Start wait group wait for job to process
	//Wait group will be stored in job, when job finish, it will call Done()
	//This is how we *hang* the request, wait for the job to be finished before returning data to user
	var wg sync.WaitGroup
	wg.Add(1)

	h.dspRequestService.SetEnv(ctx.GetRequestEnv())
	go func() {
		rtbRequest := rtb.RtbRequest{BidRequest: &bidResquest}
		rtbApiCallJob = job.DSPCallJob{
			SSP:           ssp,
			RequestData:   rtbRequest,
			Waiter:        &wg,
			RequestCaller: h.requestCaller,
			DSPService:    h.dspRequestService,
			DSPLogService: h.dspLogService,
			Config:        h.systemConfig.Rtb,
			RequestEnv:    ctx.GetRequestEnv(),
		}
		worker_pool.RtbApiCallJobQueue <- &rtbApiCallJob
	}()

	//Waiting
	wg.Wait()

	logger.ShowLog(ctx.GetRequestEnv().GetAppLogs())

	//When wait group all done, mean the rtb api call job done
	//We return data to user
	if rtbApiCallJob.BidResult == nil {
		var errResult struct {
			Errors []string `json:"errors"`
		}

		for _, e := range rtbApiCallJob.ErrorList {
			errResult.Errors = append(errResult.Errors, e.Error())
		}
		return response.Ok(errResult)
	} else {
		//Check min_floor
		if ssp.MinFloor > *rtbApiCallJob.BidResult.HighestBid.Price {
			return response.NoContent()
		}

		newNurl, err := h.callbackService.BuildSSPNUrl(rtbApiCallJob.BidResult, ssp)
		if err != nil {
			return response.ErrorResponse(err, http.StatusUnprocessableEntity)
		}
		bidResponse := rtbApiCallJob.BidResult.ToSSPResponse(ssp.ProfitMargin, newNurl)

		logger.Info(fmt.Sprintf("Profit margin is %v and bid response is", ssp.ProfitMargin))
		logger.Info(bidResponse.String())

		h.sspLogService.LogAfterRequest(ssp)

		return response.Ok(bidResponse)
	}
}

func (h RtbHandler) HandleSSPNurl(ctx tool.AppContext) response.ApiResponse {
	token := chi.URLParam(ctx.GetRequest(), "token")
	if token == "" {
		return response.ErrorResponse(
			errors.New("Win request does not have token"),
			http.StatusUnprocessableEntity,
		)
	}

	winConfirmPayload := payload.SSPWinConfirmPayload{Token: token}
	payloadJson, err := json.Marshal(winConfirmPayload)
	if err != nil {
		return response.ErrorResponse(err, http.StatusUnprocessableEntity)
	}

	err = h.messagePublisher.Publish(payloadJson, h.systemConfig.MessageQueue.WinConfirmQueue)
	if err != nil {
		return response.ErrorResponse(err, http.StatusInternalServerError)
	}

	return response.NoContent()
}
