package usecase

import (
	"context"

	apperrors "admin-pro/pkg/errors"
	"github.com/google/uuid"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

type MenuUsecase interface {
	GetMenuList(ctx context.Context) ([]*entity.Menu, error)
	GetMenu(ctx context.Context, id string) (*entity.Menu, error)
	CreateMenu(ctx context.Context, menu *entity.Menu) error
	UpdateMenu(ctx context.Context, menu *entity.Menu) error
	DeleteMenu(ctx context.Context, id string) error
	GetMenusByRoleNames(ctx context.Context, roleNames []string) ([]*entity.Menu, error)
}

type menuUsecase struct {
	menuRepo repository.MenuRepository
}

func NewMenuUsecase(menuRepo repository.MenuRepository) MenuUsecase {
	return &menuUsecase{
		menuRepo: menuRepo,
	}
}

func (u *menuUsecase) GetMenuList(ctx context.Context) ([]*entity.Menu, error) {
	return u.menuRepo.GetAllMenus(ctx)
}

func (u *menuUsecase) GetMenu(ctx context.Context, id string) (*entity.Menu, error) {
	menu, err := u.menuRepo.GetMenu(ctx, id)
	if err != nil {
		return nil, apperrors.Wrap(err, "failed to get menu")
	}
	if menu == nil {
		return nil, apperrors.ErrNotFound
	}
	return menu, nil
}

func (u *menuUsecase) CreateMenu(ctx context.Context, menu *entity.Menu) error {
	menu.ID = uuid.NewString()
	return u.menuRepo.CreateMenu(ctx, menu)
}

func (u *menuUsecase) UpdateMenu(ctx context.Context, menu *entity.Menu) error {
	exist, err := u.menuRepo.GetMenu(ctx, menu.ID)
	if err != nil {
		return apperrors.Wrap(err, "failed to get menu")
	}
	if exist == nil {
		return apperrors.ErrNotFound
	}

	return u.menuRepo.UpdateMenu(ctx, menu)
}

func (u *menuUsecase) DeleteMenu(ctx context.Context, id string) error {
	exist, err := u.menuRepo.GetMenu(ctx, id)
	if err != nil {
		return apperrors.Wrap(err, "failed to get menu")
	}
	if exist == nil {
		return apperrors.ErrNotFound
	}

	return u.menuRepo.DeleteMenu(ctx, id)
}

func (u *menuUsecase) GetMenusByRoleNames(ctx context.Context, roleNames []string) ([]*entity.Menu, error) {
	return u.menuRepo.GetMenusByRoleNames(ctx, roleNames)
}
