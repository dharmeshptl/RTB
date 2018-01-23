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

type SSPHandler struct {
	sspRepo    *repository.SSPRepository
	rtbApiRepo *repository.DSPRepository
}

func NewSSPHandler(
	sspRepo *repository.SSPRepository,
	dspRepo *repository.DSPRepository,
) SSPHandler {
	return SSPHandler{
		sspRepo,
		dspRepo,
	}
}

func (h SSPHandler) Get(ctx tool.AppContext) response.ApiResponse {
	ssp := middleware.GetSSPFromContext(ctx.GetRequest().Context())

	return response.Ok(ssp)
}

func (h SSPHandler) Update(ctx tool.AppContext) response.ApiResponse {
	ssp := middleware.GetSSPFromContext(ctx.GetRequest().Context())
	data := &payload.SSPUpdatePayload{SSP: ssp}
	if err := render.Bind(ctx.GetRequest(), data); err != nil {
		return response.ErrorResponse(err, http.StatusUnprocessableEntity)
	}

	ssp = data.SSP

	err := ctx.Validate(ssp)
	if err != nil {
		return response.ValidationError(err)
	}

	err = h.sspRepo.Update(ssp)
	if err != nil {
		return response.ErrorResponse(err, http.StatusInternalServerError)
	}

	return response.NoContent()
}

func (h SSPHandler) Create(ctx tool.AppContext) response.ApiResponse {
	data := &payload.SSPCreatePayload{}
	if err := render.Bind(ctx.GetRequest(), data); err != nil {
		return response.ErrorResponse(err, http.StatusUnprocessableEntity)
	}

	ssp := data.SSP
	err := ctx.Validate(ssp)
	if err != nil {
		return response.ValidationError(err)
	}

	err = h.sspRepo.Create(ssp)
	if err != nil {
		return response.ErrorResponse(err, http.StatusInternalServerError)
	}

	return response.NoContent()
}

func (h SSPHandler) Delete(ctx tool.AppContext) response.ApiResponse {
	ssp := middleware.GetSSPFromContext(ctx.GetRequest().Context())

	err := h.sspRepo.Delete(ssp)
	if err != nil {
		return response.ErrorResponse(err, http.StatusInternalServerError)
	}

	return response.NoContent()
}
