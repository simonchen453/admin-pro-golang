package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

type ConfigUsecase interface {
	GetConfigList(ctx context.Context) ([]*entity.Config, error)
	GetConfig(ctx context.Context, id string) (*entity.Config, error)
	GetConfigByKey(ctx context.Context, key string) (*entity.Config, error)
	CreateConfig(ctx context.Context, config *entity.Config) error
	UpdateConfig(ctx context.Context, config *entity.Config) error
	DeleteConfig(ctx context.Context, id string) error
}

type configUsecase struct {
	configRepo repository.ConfigRepository
}

func NewConfigUsecase(configRepo repository.ConfigRepository) ConfigUsecase {
	return &configUsecase{
		configRepo: configRepo,
	}
}

func (u *configUsecase) GetConfigList(ctx context.Context) ([]*entity.Config, error) {
	return u.configRepo.GetList(ctx)
}

func (u *configUsecase) GetConfig(ctx context.Context, id string) (*entity.Config, error) {
	return u.configRepo.GetByID(ctx, id)
}

func (u *configUsecase) GetConfigByKey(ctx context.Context, key string) (*entity.Config, error) {
	return u.configRepo.GetByKey(ctx, key)
}

func (u *configUsecase) CreateConfig(ctx context.Context, config *entity.Config) error {
	unique, err := u.configRepo.CheckKeyUnique(ctx, config.Key)
	if err != nil {
		return err
	}
	if !unique {
		return errors.New("config key already exists")
	}

	config.ID = uuid.NewString()
	config.CreatedDate = time.Now()
	// Default type? Let's assume N if not provided, or logic in frontend
	return u.configRepo.Create(ctx, config)
}

func (u *configUsecase) UpdateConfig(ctx context.Context, config *entity.Config) error {
	exist, err := u.configRepo.GetByID(ctx, config.ID)
	if err != nil {
		return err
	}
	if exist == nil {
		return errors.New("config not found")
	}
	config.UpdatedDate = time.Now()
	return u.configRepo.Update(ctx, config)
}

func (u *configUsecase) DeleteConfig(ctx context.Context, id string) error {
	return u.configRepo.Delete(ctx, id)
}
