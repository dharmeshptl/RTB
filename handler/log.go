package handler

import (
	"go_rtb/internal/handler/filter"
	"go_rtb/internal/handler/response"
	"go_rtb/internal/repository"
	"go_rtb/internal/tool"
	"go_rtb/internal/tool/helper"
	"net/http"
)

type LogHandler struct {
	sspLogRepo *repository.SSPLogsRepository
	dspLogRepo *repository.DSPLogsRepository
	statRepo   *repository.StatRepository
}

func NewLogHandler(
	sspLogRepo *repository.SSPLogsRepository,
	dspLogRepo *repository.DSPLogsRepository,
	statRepo *repository.StatRepository,
) LogHandler {
	return LogHandler{
		sspLogRepo,
		dspLogRepo,
		statRepo,
	}
}

func (h LogHandler) GetSSPLog(ctx tool.AppContext) response.ApiResponse {
	f := filter.NewSSPLogFilter()
	if err := helper.DecodeURLParam(f, ctx.GetRequest()); err != nil {
		return response.ErrorResponse(err, http.StatusUnprocessableEntity)
	}
	if f.StatHour == 0 {
		f.StatHour = helper.GetCurrentStatHour()
	}

	sspLog, err := h.sspLogRepo.FindFilter(f)
	if err != nil {
		return response.RecordNotFoundError(err, "Can not find log for ssp at requested time")
	}

	return response.Ok(sspLog)
}

func (h LogHandler) GetDSPLog(ctx tool.AppContext) response.ApiResponse {
	f := filter.NewDSPLogFilter()
	if err := helper.DecodeURLParam(f, ctx.GetRequest()); err != nil {
		return response.ErrorResponse(err, http.StatusUnprocessableEntity)
	}
	if f.StatHour == 0 {
		f.StatHour = helper.GetCurrentStatHour()
	}

	dspLog, err := h.dspLogRepo.FindFilter(f)
	if err != nil {
		return response.RecordNotFoundError(err, "Can not find log for dsp at requested time")
	}

	return response.Ok(dspLog)
}

func (h LogHandler) GetStat(ctx tool.AppContext) response.ApiResponse {
	f := filter.NewStatFilter()
	if err := helper.DecodeURLParam(f, ctx.GetRequest()); err != nil {
		return response.ErrorResponse(err, http.StatusUnprocessableEntity)
	}
	if f.StatHour == 0 {
		f.StatHour = helper.GetCurrentStatHour()
	}

	stat, err := h.statRepo.FindFilter(f)
	if err != nil {
		return response.RecordNotFoundError(err, "Can not find stat log at requested time")
	}

	return response.Ok(stat)
}
