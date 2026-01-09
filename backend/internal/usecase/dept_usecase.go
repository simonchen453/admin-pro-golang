package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
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
		return err
	}
	if !isUnique {
		return errors.New("department name already exists in this level")
	}
	
	dept.ID = uuid.NewString() // Generate UUID
	dept.CreatedDate = time.Now()
	dept.Status = "active" // Default status
	return u.deptRepo.Create(ctx, dept)
}

func (u *deptUsecase) UpdateDept(ctx context.Context, dept *entity.Dept) error {
	existing, err := u.deptRepo.GetByID(ctx, dept.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("department not found")
	}

	dept.UpdatedDate = time.Now()
	// Merge updates (simplified)
	return u.deptRepo.Update(ctx, dept)
}

func (u *deptUsecase) DeleteDept(ctx context.Context, id string) error {
	hasChild, err := u.deptRepo.HasChildByDeptID(ctx, id)
	if err != nil {
		return err
	}
	if hasChild {
		return errors.New("cannot delete department with sub-departments")
	}
	return u.deptRepo.Delete(ctx, id)
}
