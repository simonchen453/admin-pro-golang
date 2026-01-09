package mysql

import (
	"context"
	"errors"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
	"admin-pro/pkg/constants"
	"gorm.io/gorm"
)

type deptRepository struct {
	db *gorm.DB
}

func NewDeptRepository(db *gorm.DB) repository.DeptRepository {
	return &deptRepository{db: db}
}

func (r *deptRepository) GetList(ctx context.Context) ([]*entity.Dept, error) {
	var depts []*entity.Dept
	err := r.db.WithContext(ctx).
		Where("col_status = ?", constants.StatusActive).
		Order("col_order_num ASC").
		Find(&depts).Error
	return depts, err
}

func (r *deptRepository) GetByID(ctx context.Context, id string) (*entity.Dept, error) {
	var dept entity.Dept
	err := r.db.WithContext(ctx).Where("col_id = ?", id).First(&dept).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &dept, nil
}

func (r *deptRepository) Create(ctx context.Context, dept *entity.Dept) error {
	return r.db.WithContext(ctx).Create(dept).Error
}

func (r *deptRepository) Update(ctx context.Context, dept *entity.Dept) error {
	return r.db.WithContext(ctx).Save(dept).Error
}

func (r *deptRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("col_id = ?", id).Delete(&entity.Dept{}).Error
}

func (r *deptRepository) CheckDeptNameUnique(ctx context.Context, parentID string, deptName string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Dept{}).
		Where("col_parent_id = ? AND col_name = ?", parentID, deptName).
		Count(&count).Error
	return count == 0, err
}

func (r *deptRepository) HasChildByDeptID(ctx context.Context, deptID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Dept{}).
		Where("col_parent_id = ?", deptID).
		Count(&count).Error
	return count > 0, err
}
