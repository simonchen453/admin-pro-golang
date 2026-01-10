package usecase

import (
	"context"
	"errors"

	apperrors "admin-pro/pkg/errors"
	"github.com/google/uuid"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

type RoleUsecase interface {
	GetRoleList(ctx context.Context) ([]*entity.Role, error)
	GetRole(ctx context.Context, id string) (*entity.Role, error)
	CreateRole(ctx context.Context, role *entity.Role) error
	UpdateRole(ctx context.Context, role *entity.Role) error
	DeleteRole(ctx context.Context, id string) error
}

type roleUsecase struct {
	roleRepo repository.RoleRepository
}

func NewRoleUsecase(roleRepo repository.RoleRepository) RoleUsecase {
	return &roleUsecase{
		roleRepo: roleRepo,
	}
}

func (u *roleUsecase) GetRoleList(ctx context.Context) ([]*entity.Role, error) {
	return u.roleRepo.GetRoleList(ctx)
}

func (u *roleUsecase) GetRole(ctx context.Context, id string) (*entity.Role, error) {
	role, err := u.roleRepo.GetRole(ctx, id)
	if err != nil {
		return nil, apperrors.Wrap(err, "failed to get role")
	}
	if role == nil {
		return nil, apperrors.ErrNotFound
	}
	return role, nil
}

func (u *roleUsecase) CreateRole(ctx context.Context, role *entity.Role) error {
	// 检查角色名称是否已存在
	existingRoles, err := u.roleRepo.GetRoleList(ctx)
	if err != nil {
		return apperrors.Wrap(err, "failed to check existing roles")
	}
	for _, existingRole := range existingRoles {
		if existingRole.Name == role.Name {
			return apperrors.ErrAlreadyExists
		}
	}

	role.ID = uuid.NewString()
	return u.roleRepo.CreateRole(ctx, role)
}

func (u *roleUsecase) UpdateRole(ctx context.Context, role *entity.Role) error {
	exist, err := u.roleRepo.GetRole(ctx, role.ID)
	if err != nil {
		return apperrors.Wrap(err, "failed to get role")
	}
	if exist == nil {
		return apperrors.ErrNotFound
	}

	// 检查是否为系统角色
	if exist.IsSystem != nil && *exist.IsSystem {
		return errors.New("cannot update system role")
	}

	return u.roleRepo.UpdateRole(ctx, role)
}

func (u *roleUsecase) DeleteRole(ctx context.Context, id string) error {
	exist, err := u.roleRepo.GetRole(ctx, id)
	if err != nil {
		return apperrors.Wrap(err, "failed to get role")
	}
	if exist == nil {
		return apperrors.ErrNotFound
	}

	// 检查是否为系统角色
	if exist.IsSystem != nil && *exist.IsSystem {
		return errors.New("cannot delete system role")
	}

	return u.roleRepo.DeleteRole(ctx, id)
}
