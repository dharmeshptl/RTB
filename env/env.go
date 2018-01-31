package env

import (
	"context"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"go_rtb/internal/tool/logger"
	"gopkg.in/go-playground/validator.v9"
)

type Env struct {
	ctx         context.Context
	processLogs map[string][]logger.AppLog
	v           *validator.Validate
}

func NewEnv(ctx context.Context, v *validator.Validate) *Env {
	if ctx == nil {
		ctx = context.Background()
	}

	processLogs := make(map[string][]logger.AppLog)

	return &Env{ctx, processLogs, v}
}

func (e *Env) GetContextId() string {
	return chiMiddleware.GetReqID(e.ctx)
}

func (e *Env) AddAppLog(appLog logger.AppLog) {
	key := string(appLog.GetLogType())
	e.processLogs[key] = append(e.processLogs[key], appLog)
}

func (e *Env) GetValidator() *validator.Validate {
	return e.v
}

func (e *Env) GetAppLogs() map[string][]logger.AppLog {
	return e.processLogs
}
