package mysql

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

type genRepository struct {
	db *gorm.DB
}

func NewGenRepository(db *gorm.DB) repository.GenRepository {
	return &genRepository{db: db}
}

func (r *genRepository) GetTableList(ctx context.Context, tableName string, pageSize, pageNo int) ([]*entity.TableInfo, int64, error) {
	var list []*entity.TableInfo
	var count int64

	// Count
	sqlCounts := "SELECT count(*) FROM information_schema.tables WHERE table_schema = (SELECT DATABASE())"
	if tableName != "" {
		sqlCounts += fmt.Sprintf(" AND table_name LIKE '%%%s%%'", tableName)
	}
	if err := r.db.Raw(sqlCounts).Scan(&count).Error; err != nil {
		return nil, 0, err
	}

	// List
	offset := (pageNo - 1) * pageSize
	sqlSelect := "SELECT table_name, table_comment, create_time FROM information_schema.tables WHERE table_schema = (SELECT DATABASE())"
	if tableName != "" {
		sqlSelect += fmt.Sprintf(" AND table_name LIKE '%%%s%%'", tableName)
	}
	sqlSelect += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)

	if err := r.db.WithContext(ctx).Raw(sqlSelect).Scan(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

func (r *genRepository) GetTableColumns(ctx context.Context, tableName string) ([]*entity.ColumnInfo, error) {
	var list []*entity.ColumnInfo
	sql := "SELECT column_name, data_type, column_comment, column_key, extra FROM information_schema.columns WHERE table_schema = (SELECT DATABASE()) AND table_name = ?"
	err := r.db.WithContext(ctx).Raw(sql, tableName).Scan(&list).Error
	return list, err
}
