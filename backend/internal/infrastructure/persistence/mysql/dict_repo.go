package mysql

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

type dictRepository struct {
	db *gorm.DB
}

func NewDictRepository(db *gorm.DB) repository.DictRepository {
	return &dictRepository{db: db}
}

// --- DictType ---

func (r *dictRepository) GetDictTypeList(ctx context.Context) ([]*entity.DictType, error) {
	var list []*entity.DictType
	err := r.db.WithContext(ctx).Find(&list).Error
	return list, err
}

func (r *dictRepository) GetDictType(ctx context.Context, id string) (*entity.DictType, error) {
	var dt entity.DictType
	err := r.db.WithContext(ctx).Where("col_id = ?", id).First(&dt).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &dt, err
}

func (r *dictRepository) CreateDictType(ctx context.Context, dt *entity.DictType) error {
	return r.db.WithContext(ctx).Create(dt).Error
}

func (r *dictRepository) UpdateDictType(ctx context.Context, dt *entity.DictType) error {
	return r.db.WithContext(ctx).Save(dt).Error
}

func (r *dictRepository) DeleteDictType(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("col_id = ?", id).Delete(&entity.DictType{}).Error
}

func (r *dictRepository) CheckDictTypeUnique(ctx context.Context, typeName string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.DictType{}).
		Where("col_type = ?", typeName).
		Count(&count).Error
	return count == 0, err
}

// --- DictData ---

func (r *dictRepository) GetDictDataList(ctx context.Context, dictType string) ([]*entity.DictData, error) {
	var list []*entity.DictData
	db := r.db.WithContext(ctx).Order("col_sort ASC")
	if dictType != "" {
		db = db.Where("col_dict_type = ?", dictType)
	}
	err := db.Find(&list).Error
	return list, err
}

func (r *dictRepository) GetDictData(ctx context.Context, id string) (*entity.DictData, error) {
	var dd entity.DictData
	err := r.db.WithContext(ctx).Where("col_id = ?", id).First(&dd).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &dd, err
}

func (r *dictRepository) CreateDictData(ctx context.Context, dd *entity.DictData) error {
	return r.db.WithContext(ctx).Create(dd).Error
}

func (r *dictRepository) UpdateDictData(ctx context.Context, dd *entity.DictData) error {
	return r.db.WithContext(ctx).Save(dd).Error
}

func (r *dictRepository) DeleteDictData(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("col_id = ?", id).Delete(&entity.DictData{}).Error
}
