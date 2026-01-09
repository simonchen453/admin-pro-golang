package mysql

import (
	"context"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
	"gorm.io/gorm"
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

	// Count - 使用参数化查询防止 SQL 注入
	sqlCount := "SELECT count(*) FROM information_schema.tables WHERE table_schema = (SELECT DATABASE())"
	var args []interface{}

	if tableName != "" {
		sqlCount += " AND table_name LIKE ?"
		args = append(args, "%"+tableName+"%")
	}

	if err := r.db.WithContext(ctx).Raw(sqlCount, args...).Scan(&count).Error; err != nil {
		return nil, 0, err
	}

	// List - 使用参数化查询
	offset := (pageNo - 1) * pageSize
	sqlSelect := "SELECT table_name, table_comment, create_time FROM information_schema.tables WHERE table_schema = (SELECT DATABASE())"
	var selectArgs []interface{}

	if tableName != "" {
		sqlSelect += " AND table_name LIKE ?"
		selectArgs = append(selectArgs, "%"+tableName+"%")
	}
	sqlSelect += " LIMIT ? OFFSET ?"
	selectArgs = append(selectArgs, pageSize, offset)

	if err := r.db.WithContext(ctx).Raw(sqlSelect, selectArgs...).Scan(&list).Error; err != nil {
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
