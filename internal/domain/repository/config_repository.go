package repository

import (
	"admin-pro/internal/domain/entity"
	"context"
)

type ConfigRepository interface {
	GetList(ctx context.Context) ([]*entity.Config, error)
	GetByID(ctx context.Context, id string) (*entity.Config, error)
	GetByKey(ctx context.Context, key string) (*entity.Config, error)
	Create(ctx context.Context, config *entity.Config) error
	Update(ctx context.Context, config *entity.Config) error
	Delete(ctx context.Context, id string) error
	CheckKeyUnique(ctx context.Context, key string) (bool, error)
}
