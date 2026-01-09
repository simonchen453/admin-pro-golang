package usecase

import (
	"context"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

type LogUsecase interface {
	GetLoginLogList(ctx context.Context) ([]*entity.LoginLog, error)
	DeleteLoginLog(ctx context.Context, id string) error
	CleanLoginLog(ctx context.Context) error

	GetOperLogList(ctx context.Context) ([]*entity.OperLog, error)
	DeleteOperLog(ctx context.Context, id string) error
	CleanOperLog(ctx context.Context) error
}

type logUsecase struct {
	logRepo repository.LogRepository
}

func NewLogUsecase(logRepo repository.LogRepository) LogUsecase {
	return &logUsecase{
		logRepo: logRepo,
	}
}

func (u *logUsecase) GetLoginLogList(ctx context.Context) ([]*entity.LoginLog, error) {
	return u.logRepo.GetLoginLogList(ctx)
}

func (u *logUsecase) DeleteLoginLog(ctx context.Context, id string) error {
	return u.logRepo.DeleteLoginLog(ctx, id)
}

func (u *logUsecase) CleanLoginLog(ctx context.Context) error {
	return u.logRepo.CleanLoginLog(ctx)
}

func (u *logUsecase) GetOperLogList(ctx context.Context) ([]*entity.OperLog, error) {
	return u.logRepo.GetOperLogList(ctx)
}

func (u *logUsecase) DeleteOperLog(ctx context.Context, id string) error {
	return u.logRepo.DeleteOperLog(ctx, id)
}

func (u *logUsecase) CleanOperLog(ctx context.Context) error {
	return u.logRepo.CleanOperLog(ctx)
}
