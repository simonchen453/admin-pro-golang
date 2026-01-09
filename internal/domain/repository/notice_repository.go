package repository

import (
	"admin-pro/internal/domain/entity"
	"context"
)

type NoticeRepository interface {
	GetList(ctx context.Context) ([]*entity.Notice, error)
	GetByID(ctx context.Context, id string) (*entity.Notice, error)
	Create(ctx context.Context, notice *entity.Notice) error
	Update(ctx context.Context, notice *entity.Notice) error
	Delete(ctx context.Context, id string) error
}
