package handler

import (
	"github.com/go-chi/render"
	"go_rtb/internal/handler/payload"
	"go_rtb/internal/handler/response"
	"go_rtb/internal/repository"
	"go_rtb/internal/router/middleware"
	"go_rtb/internal/tool"
	"net/http"
)

type DSPHandler struct {
	sspRepo *repository.SSPRepository
	apiRepo *repository.DSPRepository
}

func NewDSPApiHandler(
	sspRepo *repository.SSPRepository,
	dspRepo *repository.DSPRepository) DSPHandler {
	return DSPHandler{
		sspRepo,
		dspRepo,
	}
}

func (h DSPHandler) Get(ctx tool.AppContext) response.ApiResponse {
	rtbApi := middleware.GetDSPFromContext(ctx.GetRequest().Context())

	return response.Ok(rtbApi)
}

func (h DSPHandler) Update(ctx tool.AppContext) response.ApiResponse {
	dsp := middleware.GetDSPFromContext(ctx.GetRequest().Context())
	data := &payload.DSPUpdatePayload{DSP: dsp}
	if err := render.Bind(ctx.GetRequest(), data); err != nil {
		return response.ErrorResponse(err, http.StatusUnprocessableEntity)
	}

	dsp = data.DSP

	err := ctx.Validate(dsp)
	if err != nil {
		return response.ValidationError(err)
	}

	err = h.apiRepo.Update(dsp)

	if err != nil {
		return response.ErrorResponse(err, http.StatusInternalServerError)
	}

	return response.NoContent()
}

func (h DSPHandler) Create(ctx tool.AppContext) response.ApiResponse {
	data := &payload.DSPCreatePayload{}
	if err := render.Bind(ctx.GetRequest(), data); err != nil {
		return response.ErrorResponse(err, http.StatusUnprocessableEntity)
	}

	dsp := data.DSP
	err := ctx.Validate(dsp)
	if err != nil {
		return response.ValidationError(err)
	}

	err = h.apiRepo.Create(dsp)

	if err != nil {
		return response.ErrorResponse(err, http.StatusInternalServerError)
	}

	return response.NoContent()
}

func (h DSPHandler) Delete(ctx tool.AppContext) response.ApiResponse {
	rtbApi := middleware.GetDSPFromContext(ctx.GetRequest().Context())

	err := h.apiRepo.Delete(rtbApi)
	if err != nil {
		return response.ErrorResponse(err, http.StatusInternalServerError)
	}

	return response.NoContent()
}
