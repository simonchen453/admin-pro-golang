package mysql

import (
	"context"

	"gorm.io/gorm"
	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

type monitorRepository struct {
	db *gorm.DB
}

func NewMonitorRepository(db *gorm.DB) repository.MonitorRepository {
	return &monitorRepository{db: db}
}

func (r *monitorRepository) GetSessionList(ctx context.Context) ([]*entity.Session, error) {
	var list []*entity.Session
	err := r.db.WithContext(ctx).Order("col_created_date DESC").Find(&list).Error
	return list, err
}

func (r *monitorRepository) DeleteSession(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("col_id = ?", id).Delete(&entity.Session{}).Error
}

func (r *monitorRepository) CreateSession(ctx context.Context, session *entity.Session) error {
	return r.db.WithContext(ctx).Create(session).Error
}
