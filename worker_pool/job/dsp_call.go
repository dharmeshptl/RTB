package job

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"go_rtb/internal/configuration"
	"go_rtb/internal/env"
	"go_rtb/internal/model"
	"go_rtb/internal/protocol_buffer"
	"go_rtb/internal/rtb"
	"go_rtb/internal/service"
	"go_rtb/internal/tool/helper"
	"go_rtb/internal/tool/logger"
	"go_rtb/internal/transport"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type key int

var ErrorNoSuitableDSPForSSPRequest = errors.New("Can not find any suitable dsp for ssp request")

type DSPCallJob struct {
	//Store ssp does rtb request
	SSP *model.SSP

	//Request data ssp call us
	RequestData rtb.RtbRequest

	//We have wait group here, waiting for job to be finished
	Waiter *sync.WaitGroup

	BidResult *rtb.BidResult

	ErrorList []error

	RequestCaller *transport.HttpCaller

	DSPService *service.DSPRequestService

	DSPLogService *service.DSPLogService

	Config configuration.RtbConfig

	RequestEnv *env.Env
}

//Process job, job has all data to be self-processed
func (j *DSPCallJob) Process() {
	dspToCalledList, errs := j.DSPService.GetDSPToCall(j.SSP, j.RequestData.BidRequest)
	for _, err := range errs {
		j.ErrorList = append(j.ErrorList, err)
	}

	rtbApiCount := len(dspToCalledList)
	if rtbApiCount == 0 {
		j.ErrorList = append(j.ErrorList, ErrorNoSuitableDSPForSSPRequest)
	} else {
		//Must be buffered channel here
		resultc := make(chan rtb.RtbApiCallResponse, rtbApiCount)
		errc := make(chan error, rtbApiCount)

		go func() {
			//Must wait for all rtb api call completed before process response
			var wg sync.WaitGroup
			wg.Add(rtbApiCount)

			for _, dsp := range dspToCalledList {
				go func(dsp *model.DSP) {
					callResult, err := j.callApi(dsp)
					resultc <- callResult
					errc <- err
					wg.Done()
				}(dsp)
			}

			wg.Wait()
			//Close result channel means all api call complete
			close(resultc)
			close(errc)
		}()

		//Process result of all call
		var highestBidResp *protocol_buffer.BidResponse
		var highestBid *protocol_buffer.BidResponse_SeatBid_Bid
		var highestPrice float64 = 0
		var highestResponseString string = ""
		var highestSeatBidPos int = 0
		var highestBidPos int = 0
		var dsp *model.DSP

		for apiResp := range resultc {
			if apiResp.Success() == true {
				for i, seatBid := range apiResp.BidResponse.Seatbid {
					for j, bid := range seatBid.Bid {
						price := bid.Price

						if highestPrice < *price {
							highestPrice = *price
							highestBidResp = &apiResp.BidResponse
							highestBid = bid
							highestResponseString = apiResp.ResponseBody
							highestSeatBidPos = i
							highestBidPos = j
							dsp = apiResp.GetDSP()
						}
					}
				}
			} else {
				if apiResp.GetError() != nil {
					j.ErrorList = append(j.ErrorList, apiResp.GetError())
				}
			}

		}

		bidResult := rtb.BidResult{}
		bidResult.HighestBidResp = highestBidResp
		bidResult.HighestBid = highestBid
		bidResult.HighestApiResponseBody = highestResponseString
		bidResult.HighestSeatBidPos = highestSeatBidPos
		bidResult.HighestBidPos = highestBidPos
		bidResult.DSP = dsp
		j.BidResult = &bidResult
		for e := range errc {
			if e != nil {
				j.ErrorList = append(j.ErrorList, e)
			}
		}
	}
	//This is shit
	//Waiting for all client's api to be call
	j.Waiter.Done()
}

//For for each rtb api call request, and build result data
func (j *DSPCallJob) callApi(dsp *model.DSP) (rtb.RtbApiCallResponse, error) {
	j.DSPLogService.LogPreRequest(dsp)
	// ctx is the Context for this handler. Calling cancel closes the
	// ctx.Done channel, which is the cancellation signal for requests
	// started by this handler.
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithTimeout(
		context.Background(),
		time.Millisecond*time.Duration(j.Config.DspRequestTimeOut),
	)
	defer cancel()

	postData := []byte(j.RequestData.BidRequest.String())
	req, _ := http.NewRequest(http.MethodPost, dsp.EndpointURL, bytes.NewBuffer(postData))
	req.Header.Set("Content-Type", "application/json")

	var apiResponse rtb.RtbApiCallResponse
	bidResponse := protocol_buffer.BidResponse{}
	err := j.RequestCaller.HttpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			apiResponse = rtb.NewApiResponse(
				dsp,
				j.RequestData.BidRequest.String(),
				resp.StatusCode,
				"",
				err,
				"",
				false,
			)

			j.RequestEnv.AddAppLog(
				logger.NewAppLog(
					logger.DspUrlCallLog,
					err,
					helper.GetCurrentTime(),
					fmt.Sprintf(
						"Request URL: %s - with Response: %s",
						dsp.EndpointURL,
						j.RequestData.BidRequest.String(),
					),
				),
			)

			return err
		}
		defer resp.Body.Close()

		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		ioReaderForJson := bytes.NewReader(bodyBytes)
		bodyString := string(bodyBytes)
		err = jsonpb.Unmarshal(ioReaderForJson, &bidResponse)
		if err != nil {
			apiResponse = rtb.NewApiResponse(
				dsp,
				j.RequestData.BidRequest.String(),
				resp.StatusCode,
				bodyString,
				err,
				"",
				false,
			)

			j.RequestEnv.AddAppLog(
				logger.NewAppLog(
					logger.RTBResponseCallLog,
					err,
					helper.GetCurrentTime(),
					apiResponse.String(),
				),
			)

			return err
		}

		j.DSPLogService.LogAfterRequest(dsp)

		apiResponse = rtb.NewApiResponse(
			dsp,
			j.RequestData.BidRequest.String(),
			resp.StatusCode,
			bodyString,
			nil,
			bidResponse.String(),
			true,
		)

		j.RequestEnv.AddAppLog(
			logger.NewAppLog(
				logger.RTBResponseCallLog,
				nil,
				helper.GetCurrentTime(),
				apiResponse.String(),
			),
		)

		apiResponse.BidResponse = bidResponse
		apiResponse.ResponseBody = bodyString
		return nil
	})

	return apiResponse, err
}
