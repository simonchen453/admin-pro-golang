package repository

import (
	"context"
	"admin-pro/internal/domain/entity"
)

type DeptRepository interface {
	GetList(ctx context.Context) ([]*entity.Dept, error)
	GetByID(ctx context.Context, id string) (*entity.Dept, error)
	Create(ctx context.Context, dept *entity.Dept) error
	Update(ctx context.Context, dept *entity.Dept) error
	Delete(ctx context.Context, id string) error
	CheckDeptNameUnique(ctx context.Context, parentID string, deptName string) (bool, error)
	HasChildByDeptID(ctx context.Context, deptID string) (bool, error)
}
