package repository

import (
	"context"
	"admin-pro/internal/domain/entity"
)

type GenRepository interface {
	GetTableList(ctx context.Context, tableName string, pageSize, pageNo int) ([]*entity.TableInfo, int64, error)
	GetTableColumns(ctx context.Context, tableName string) ([]*entity.ColumnInfo, error)
}
