package repository

import (
	"admin-pro/internal/domain/entity"
	"context"
)

type GenRepository interface {
	GetTableList(ctx context.Context, tableName string, pageSize, pageNo int) ([]*entity.TableInfo, int64, error)
	GetTableColumns(ctx context.Context, tableName string) ([]*entity.ColumnInfo, error)
}
