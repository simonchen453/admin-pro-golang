package mysql

import (
	"context"
	"errors"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
	"gorm.io/gorm"
)

type configRepository struct {
	db *gorm.DB
}

func NewConfigRepository(db *gorm.DB) repository.ConfigRepository {
	return &configRepository{db: db}
}

func (r *configRepository) GetList(ctx context.Context) ([]*entity.Config, error) {
	var list []*entity.Config
	err := r.db.WithContext(ctx).Find(&list).Error
	return list, err
}

func (r *configRepository) GetByID(ctx context.Context, id string) (*entity.Config, error) {
	var cfg entity.Config
	err := r.db.WithContext(ctx).Where("col_id = ?", id).First(&cfg).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &cfg, err
}

func (r *configRepository) GetByKey(ctx context.Context, key string) (*entity.Config, error) {
	var cfg entity.Config
	err := r.db.WithContext(ctx).Where("col_key = ?", key).First(&cfg).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &cfg, err
}

func (r *configRepository) Create(ctx context.Context, config *entity.Config) error {
	return r.db.WithContext(ctx).Create(config).Error
}

func (r *configRepository) Update(ctx context.Context, config *entity.Config) error {
	return r.db.WithContext(ctx).Save(config).Error
}

func (r *configRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("col_id = ?", id).Delete(&entity.Config{}).Error
}

func (r *configRepository) CheckKeyUnique(ctx context.Context, key string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Config{}).
		Where("col_key = ?", key).
		Count(&count).Error
	return count == 0, err
}
