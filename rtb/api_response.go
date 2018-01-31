package rtb

import (
	"fmt"
	"go_rtb/internal/model"
	"go_rtb/internal/protocol_buffer"
)

type RtbApiCallResponse struct {
	BidResponse  protocol_buffer.BidResponse
	ResponseBody string
	dsp          *model.DSP
	requestData  string
	responseCode int
	responseBody string
	err          error
	parseResult  string
	success      bool
}

func NewApiResponse(dsp *model.DSP,
	requestData string,
	responseCode int,
	responseBody string,
	err error,
	parseResult string,
	success bool) RtbApiCallResponse {
	return RtbApiCallResponse{
		dsp:          dsp,
		requestData:  requestData,
		responseCode: responseCode,
		responseBody: responseBody,
		err:          err,
		parseResult:  parseResult,
		success:      success,
	}
}

func (resp *RtbApiCallResponse) GetDSP() *model.DSP {
	return resp.dsp
}

func (resp *RtbApiCallResponse) String() string {
	var result string
	result = "----------------------------------\n"
	if resp.success == true {
		result += fmt.Sprintf("Request URL: %s\n", resp.dsp.EndpointURL)
		result += fmt.Sprintf("Request Data: %s\n", resp.requestData)
		result += fmt.Sprintf("Response Code: %d\n", resp.responseCode)
		result += fmt.Sprintf("Response Body: %s\n", resp.responseBody)
		result += fmt.Sprintf("Parse Result: %s\n", resp.parseResult)
	} else {
		result += fmt.Sprintf("Request URL: %s\n", resp.dsp)
		result += fmt.Sprintf("Request Data: %s\n", resp.requestData)
		result += fmt.Sprintf("Response Code: %d\n", resp.responseCode)
		result += fmt.Sprintf("Response Body: %s\n", resp.responseBody)
		result += fmt.Sprintf("Parse Error: %s\n", resp.err.Error())
	}
	result += "----------------------------------\n"

	return result
}

func (resp *RtbApiCallResponse) Success() bool {
	return resp.success
}

func (resp *RtbApiCallResponse) GetError() error {
	return resp.err
}
