package repository

import (
	"admin-pro/internal/domain/entity"
	"context"
)

type LogRepository interface {
	// Login Log
	CreateLoginLog(ctx context.Context, log *entity.LoginLog) error
	GetLoginLogList(ctx context.Context) ([]*entity.LoginLog, error)
	DeleteLoginLog(ctx context.Context, id string) error
	CleanLoginLog(ctx context.Context) error

	// Oper Log
	CreateOperLog(ctx context.Context, log *entity.OperLog) error
	GetOperLogList(ctx context.Context) ([]*entity.OperLog, error)
	DeleteOperLog(ctx context.Context, id string) error
	CleanOperLog(ctx context.Context) error
}
