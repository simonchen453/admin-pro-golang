package usecase

import (
	"context"
	"time"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
	apperrors "admin-pro/pkg/errors"
	"github.com/google/uuid"
)

type DeptUsecase interface {
	GetDeptList(ctx context.Context) ([]*entity.Dept, error)
	GetDept(ctx context.Context, id string) (*entity.Dept, error)
	CreateDept(ctx context.Context, dept *entity.Dept) error
	UpdateDept(ctx context.Context, dept *entity.Dept) error
	DeleteDept(ctx context.Context, id string) error
}

type deptUsecase struct {
	deptRepo repository.DeptRepository
}

func NewDeptUsecase(deptRepo repository.DeptRepository) DeptUsecase {
	return &deptUsecase{
		deptRepo: deptRepo,
	}
}

func (u *deptUsecase) GetDeptList(ctx context.Context) ([]*entity.Dept, error) {
	return u.deptRepo.GetList(ctx)
}

func (u *deptUsecase) GetDept(ctx context.Context, id string) (*entity.Dept, error) {
	return u.deptRepo.GetByID(ctx, id)
}

func (u *deptUsecase) CreateDept(ctx context.Context, dept *entity.Dept) error {
	isUnique, err := u.deptRepo.CheckDeptNameUnique(ctx, dept.ParentID, dept.Name)
	if err != nil {
		return apperrors.Wrap(err, "failed to check department name uniqueness")
	}
	if !isUnique {
		return apperrors.ErrAlreadyExists
	}

	dept.ID = uuid.NewString() // 生成 UUID
	dept.CreatedDate = time.Now()
	dept.Status = "active" // 默认状态
	return u.deptRepo.Create(ctx, dept)
}

func (u *deptUsecase) UpdateDept(ctx context.Context, dept *entity.Dept) error {
	existing, err := u.deptRepo.GetByID(ctx, dept.ID)
	if err != nil {
		return apperrors.Wrap(err, "failed to get department")
	}
	if existing == nil {
		return apperrors.ErrNotFound
	}

	dept.UpdatedDate = time.Now()
	// 合并更新
	return u.deptRepo.Update(ctx, dept)
}

func (u *deptUsecase) DeleteDept(ctx context.Context, id string) error {
	hasChild, err := u.deptRepo.HasChildByDeptID(ctx, id)
	if err != nil {
		return apperrors.Wrap(err, "failed to check sub-departments")
	}
	if hasChild {
		return apperrors.Wrap(apperrors.ErrInvalidOperation, "cannot delete department with sub-departments")
	}
	return u.deptRepo.Delete(ctx, id)
}
