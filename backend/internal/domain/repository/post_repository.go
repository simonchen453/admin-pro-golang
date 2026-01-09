package repository

import (
	"context"
	"admin-pro/internal/domain/entity"
)

type PostRepository interface {
	GetList(ctx context.Context) ([]*entity.Post, error)
	GetByID(ctx context.Context, id string) (*entity.Post, error)
	Create(ctx context.Context, post *entity.Post) error
	Update(ctx context.Context, post *entity.Post) error
	Delete(ctx context.Context, id string) error
	CheckPostCodeUnique(ctx context.Context, code string) (bool, error)
	CheckPostNameUnique(ctx context.Context, name string) (bool, error)
}
