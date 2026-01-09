package usecase

import (
	"context"
	"errors"
	"time"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
	"github.com/google/uuid"
)

type DictUsecase interface {
	// Type
	GetDictTypeList(ctx context.Context) ([]*entity.DictType, error)
	GetDictType(ctx context.Context, id string) (*entity.DictType, error)
	CreateDictType(ctx context.Context, dt *entity.DictType) error
	UpdateDictType(ctx context.Context, dt *entity.DictType) error
	DeleteDictType(ctx context.Context, id string) error

	// Data
	GetDictDataByType(ctx context.Context, dictType string) ([]*entity.DictData, error)
	GetDictData(ctx context.Context, id string) (*entity.DictData, error)
	CreateDictData(ctx context.Context, dd *entity.DictData) error
	UpdateDictData(ctx context.Context, dd *entity.DictData) error
	DeleteDictData(ctx context.Context, id string) error
}

type dictUsecase struct {
	dictRepo repository.DictRepository
}

func NewDictUsecase(dictRepo repository.DictRepository) DictUsecase {
	return &dictUsecase{
		dictRepo: dictRepo,
	}
}

// --- Type ---

func (u *dictUsecase) GetDictTypeList(ctx context.Context) ([]*entity.DictType, error) {
	return u.dictRepo.GetDictTypeList(ctx)
}

func (u *dictUsecase) GetDictType(ctx context.Context, id string) (*entity.DictType, error) {
	return u.dictRepo.GetDictType(ctx, id)
}

func (u *dictUsecase) CreateDictType(ctx context.Context, dt *entity.DictType) error {
	unique, err := u.dictRepo.CheckDictTypeUnique(ctx, dt.Type)
	if err != nil {
		return err
	}
	if !unique {
		return errors.New("dict type already exists")
	}

	dt.ID = uuid.NewString()
	dt.CreatedDate = time.Now()
	dt.Status = "active"
	return u.dictRepo.CreateDictType(ctx, dt)
}

func (u *dictUsecase) UpdateDictType(ctx context.Context, dt *entity.DictType) error {
	exist, err := u.dictRepo.GetDictType(ctx, dt.ID)
	if err != nil {
		return err
	}
	if exist == nil {
		return errors.New("dict type not found")
	}
	dt.UpdatedDate = time.Now()
	return u.dictRepo.UpdateDictType(ctx, dt)
}

func (u *dictUsecase) DeleteDictType(ctx context.Context, id string) error {
	return u.dictRepo.DeleteDictType(ctx, id)
}

// --- Data ---

func (u *dictUsecase) GetDictDataByType(ctx context.Context, dictType string) ([]*entity.DictData, error) {
	return u.dictRepo.GetDictDataList(ctx, dictType)
}

func (u *dictUsecase) GetDictData(ctx context.Context, id string) (*entity.DictData, error) {
	return u.dictRepo.GetDictData(ctx, id)
}

func (u *dictUsecase) CreateDictData(ctx context.Context, dd *entity.DictData) error {
	dd.ID = uuid.NewString()
	dd.CreatedDate = time.Now()
	dd.Status = "active"
	return u.dictRepo.CreateDictData(ctx, dd)
}

func (u *dictUsecase) UpdateDictData(ctx context.Context, dd *entity.DictData) error {
	exist, err := u.dictRepo.GetDictData(ctx, dd.ID)
	if err != nil {
		return err
	}
	if exist == nil {
		return errors.New("dict data not found")
	}
	dd.UpdatedDate = time.Now()
	return u.dictRepo.UpdateDictData(ctx, dd)
}

func (u *dictUsecase) DeleteDictData(ctx context.Context, id string) error {
	return u.dictRepo.DeleteDictData(ctx, id)
}
