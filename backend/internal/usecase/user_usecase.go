package usecase

import (
	"context"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

// UserUsecase 定义了用户相关的业务逻辑接口
// 位于 Usecase 层，协调 Entity 和 Repository
type UserUsecase interface {
	GetUserInfo(ctx context.Context, userID string) (*UserInfoDTO, error)    // 获取用户信息（包含权限）
	GetUserMenus(ctx context.Context, userID string) ([]*entity.Menu, error) // 获取用户菜单
}

// UserInfoDTO 是 Data Transfer Object (数据传输对象)
// 用于 Usecase 层向 Delivery 层返回组合数据，不直接暴露 Entity
type UserInfoDTO struct {
	User        *entity.User
	Roles       []string
	Permissions []string
}

type userUsecase struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
	menuRepo repository.MenuRepository
}

// NewUserUsecase 构造函数，依赖注入 Repositories
func NewUserUsecase(
	userRepo repository.UserRepository,
	roleRepo repository.RoleRepository,
	menuRepo repository.MenuRepository,
) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		roleRepo: roleRepo,
		menuRepo: menuRepo,
	}
}

// GetUserInfo 实现获取用户信息的业务逻辑
func (u *userUsecase) GetUserInfo(ctx context.Context, userID string) (*UserInfoDTO, error) {
	// 1. 查询用户基本信息
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil // 用户不存在
	}

	// 2. 查询用户拥有的角色
	roles, err := u.roleRepo.GetRolesByUserID(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	var roleNames []string
	for _, r := range roles {
		roleNames = append(roleNames, r.Name)
	}

	// 3. 根据角色查询拥有的菜单/权限
	menus, err := u.menuRepo.GetMenusByRoleNames(ctx, roleNames)
	if err != nil {
		return nil, err
	}

	var permissions []string
	for _, m := range menus {
		if m.Permission != "" {
			permissions = append(permissions, m.Permission)
		}
	}

	// 4. 组装结果返回
	return &UserInfoDTO{
		User:        user,
		Roles:       roleNames,
		Permissions: permissions,
	}, nil
}

// GetUserMenus 获取用户可访问的菜单树
func (u *userUsecase) GetUserMenus(ctx context.Context, userID string) ([]*entity.Menu, error) {
	// 1. 获取用户角色
	roles, err := u.roleRepo.GetRolesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var roleNames []string
	isSuperAdmin := false
	for _, r := range roles {
		roleNames = append(roleNames, r.Name)
		if r.Name == "SYS_SUPER_ADMIN_ROLE" || r.Name == "SYS_ADMIN_ROLE" {
			isSuperAdmin = true // 超级管理员拥有所有权限
		}
	}

	// 2. 如果是超管，直接返回所有菜单
	if isSuperAdmin {
		return u.menuRepo.GetAllMenus(ctx)
	}

	// 3. 否则根据角色筛选菜单
	return u.menuRepo.GetMenusByRoleNames(ctx, roleNames)
}
