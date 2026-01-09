package repository

import (
	"admin-pro/internal/domain/entity"
	"context"
)

// UserRepository 定义了用户相关的数据库操作接口
// 这是一个 interface，只定义“做什么”，不定义“怎么做”
// 具体的实现（如 MySQL, Redis）在 Infrastructure 层
type UserRepository interface {
	GetByLoginName(ctx context.Context, loginName string) (*entity.User, error) // 根据登录名查询用户
	GetByID(ctx context.Context, id string) (*entity.User, error)               // 根据ID查询用户
}

// RoleRepository 定义了角色相关的数据库操作接口
type RoleRepository interface {
	GetRolesByUserID(ctx context.Context, userID string) ([]*entity.Role, error) // 查询某个用户拥有的所有角色
	GetRoleList(ctx context.Context) ([]*entity.Role, error)                     // 查询所有角色
	GetRole(ctx context.Context, id string) (*entity.Role, error)                // 根据ID查询角色
	CreateRole(ctx context.Context, role *entity.Role) error                     // 创建角色
	UpdateRole(ctx context.Context, role *entity.Role) error                     // 更新角色
	DeleteRole(ctx context.Context, id string) error                             // 删除角色
}

// MenuRepository 定义了菜单相关的数据库操作接口
type MenuRepository interface {
	GetMenusByRoleNames(ctx context.Context, roleNames []string) ([]*entity.Menu, error) // 根据角色名列表查询可访问的菜单
	GetAllMenus(ctx context.Context) ([]*entity.Menu, error)                             // 查询所有菜单
	GetMenu(ctx context.Context, id string) (*entity.Menu, error)                        // 根据ID查询菜单
	CreateMenu(ctx context.Context, menu *entity.Menu) error                             // 创建菜单
	UpdateMenu(ctx context.Context, menu *entity.Menu) error                             // 更新菜单
	DeleteMenu(ctx context.Context, id string) error                                     // 删除菜单
}
