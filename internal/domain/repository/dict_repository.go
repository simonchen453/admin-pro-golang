package repository

import (
	"admin-pro/internal/domain/entity"
	"context"
)

type DictRepository interface {
	// DictType
	GetDictTypeList(ctx context.Context) ([]*entity.DictType, error)
	GetDictType(ctx context.Context, id string) (*entity.DictType, error)
	CreateDictType(ctx context.Context, dt *entity.DictType) error
	UpdateDictType(ctx context.Context, dt *entity.DictType) error
	DeleteDictType(ctx context.Context, id string) error
	CheckDictTypeUnique(ctx context.Context, typeName string) (bool, error)

	// DictData
	GetDictDataList(ctx context.Context, dictType string) ([]*entity.DictData, error)
	GetDictData(ctx context.Context, id string) (*entity.DictData, error)
	CreateDictData(ctx context.Context, dd *entity.DictData) error
	UpdateDictData(ctx context.Context, dd *entity.DictData) error
	DeleteDictData(ctx context.Context, id string) error
}
