package mysql

import (
	"context"
	"errors"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
	"gorm.io/gorm"
)

type noticeRepository struct {
	db *gorm.DB
}

func NewNoticeRepository(db *gorm.DB) repository.NoticeRepository {
	return &noticeRepository{db: db}
}

func (r *noticeRepository) GetList(ctx context.Context) ([]*entity.Notice, error) {
	var list []*entity.Notice
	err := r.db.WithContext(ctx).Order("col_created_date DESC").Find(&list).Error
	return list, err
}

func (r *noticeRepository) GetByID(ctx context.Context, id string) (*entity.Notice, error) {
	var notice entity.Notice
	err := r.db.WithContext(ctx).Where("col_id = ?", id).First(&notice).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &notice, err
}

func (r *noticeRepository) Create(ctx context.Context, notice *entity.Notice) error {
	return r.db.WithContext(ctx).Create(notice).Error
}

func (r *noticeRepository) Update(ctx context.Context, notice *entity.Notice) error {
	return r.db.WithContext(ctx).Save(notice).Error
}

func (r *noticeRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("col_id = ?", id).Delete(&entity.Notice{}).Error
}
