package mysql

import (
	"context"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
	"gorm.io/gorm"
)

type logRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) repository.LogRepository {
	return &logRepository{db: db}
}

// --- Login Log ---

func (r *logRepository) CreateLoginLog(ctx context.Context, log *entity.LoginLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *logRepository) GetLoginLogList(ctx context.Context) ([]*entity.LoginLog, error) {
	var list []*entity.LoginLog
	err := r.db.WithContext(ctx).Order("col_login_time DESC").Limit(100).Find(&list).Error // Limit default 100 for now
	return list, err
}

func (r *logRepository) DeleteLoginLog(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("col_id = ?", id).Delete(&entity.LoginLog{}).Error
}

func (r *logRepository) CleanLoginLog(ctx context.Context) error {
	return r.db.WithContext(ctx).Exec("TRUNCATE TABLE sys_login_info_tbl").Error
}

// --- Oper Log ---

func (r *logRepository) CreateOperLog(ctx context.Context, log *entity.OperLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *logRepository) GetOperLogList(ctx context.Context) ([]*entity.OperLog, error) {
	var list []*entity.OperLog
	err := r.db.WithContext(ctx).Order("col_log_date DESC").Limit(100).Find(&list).Error
	return list, err
}

func (r *logRepository) DeleteOperLog(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("col_id = ?", id).Delete(&entity.OperLog{}).Error
}

func (r *logRepository) CleanOperLog(ctx context.Context) error {
	return r.db.WithContext(ctx).Exec("TRUNCATE TABLE sys_audit_log_tbl").Error
}
