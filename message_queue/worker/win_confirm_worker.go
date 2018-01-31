package worker

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_rtb/internal/message_queue/payload"
	"go_rtb/internal/model"
	"go_rtb/internal/repository"
	"go_rtb/internal/service"
	"go_rtb/internal/tool/logger"
	"net/http"
)

type WinConfirmWorker struct {
	nurlRepo      *repository.NUrlRepository
	statService   *service.StatService
	sspLogService *service.SSPLogService
	dspLogService *service.DSPLogService
	sspRepo       *repository.SSPRepository
	dspRepo       *repository.DSPRepository
}

func NewWinConfirmWorker(
	nurlRepo *repository.NUrlRepository,
	statService *service.StatService,
	sspLogService *service.SSPLogService,
	dspLogService *service.DSPLogService,
	sspRepo *repository.SSPRepository,
	dspRepo *repository.DSPRepository,
) *WinConfirmWorker {
	return &WinConfirmWorker{
		nurlRepo,
		statService,
		sspLogService,
		dspLogService,
		sspRepo,
		dspRepo,
	}
}

func (worker *WinConfirmWorker) Process(message []byte) error {
	var winConfirmPayload payload.SSPWinConfirmPayload
	if err := json.Unmarshal(message, &winConfirmPayload); err != nil {
		return err
	}

	nurl, err := worker.nurlRepo.FindById(winConfirmPayload.Token)
	if err != nil {
		return err
	}

	if nurl.IsUsed {
		return errors.New(
			fmt.Sprintf("Nurl token [%s] was already used", nurl.Token),
		)
	}

	err = worker.nurlRepo.MarkAsUsed(nurl)
	if err != nil {
		return err
	}

	dsp, _ := worker.dspRepo.FindByID(nurl.DSPId)
	ssp, _ := worker.sspRepo.FindByID(nurl.SSPId)

	err = worker.callDSPWonUrl(nurl)
	if err != nil {
		return err
	}

	return worker.logSSPWin(nurl, ssp, dsp)
}

func (worker *WinConfirmWorker) callDSPWonUrl(nurl *model.NUrl) error {
	//TODO: retry multiple time
	resp, err := http.Get(nurl.DSPNurl)
	if err != nil {
		return err
	}

	logger.Info(
		fmt.Sprintf(
			"Satus code [%d] response when calling nurl: %s. Token: %s",
			resp.StatusCode,
			nurl.DSPNurl,
			nurl.Token,
		),
	)

	return nil
}

func (worker *WinConfirmWorker) logSSPWin(nurl *model.NUrl, ssp *model.SSP, dsp *model.DSP) error {
	worker.sspLogService.LogWinConfirm(nurl, ssp)
	worker.dspLogService.LogWinConfirm(nurl, dsp)
	worker.statService.LogWinConfirm(nurl, ssp, dsp)

	return nil
}
