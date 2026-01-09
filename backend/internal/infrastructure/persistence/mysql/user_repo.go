package mysql

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

// UserRepository 的 MySQL 实现
// 这里是真正执行 SQL 的地方，依赖 GORM
type userRepository struct {
	db *gorm.DB // 持有 GORM DB 对象
}

// NewUserRepository 构造函数
// 返回 repository.UserRepository 接口，保证依赖倒置
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

// GetByLoginName 根据登录名查询用户
// 实现接口中定义的方法
func (r *userRepository) GetByLoginName(ctx context.Context, loginName string) (*entity.User, error) {
	var user entity.User
	// GORM 查询：SELECT * FROM sys_user_tbl WHERE col_login_name = ? LIMIT 1
	if err := r.db.WithContext(ctx).Where("col_login_name = ?", loginName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 没找到不报错，返回 nil
		}
		return nil, err // 数据库错误
	}
	return &user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("col_id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) GetRolesByUserID(ctx context.Context, userID string) ([]*entity.Role, error) {
	var roles []*entity.Role
	// 多表联查：通过 sys_user_role_assign_tbl 关联表查询角色信息
	// JOIN sys_user_role_assign_tbl ura ON r.col_name = ura.col_role_name WHERE ura.col_user_id = ?
	err := r.db.WithContext(ctx).
		Table("sys_role_tbl r").
		Joins("JOIN sys_user_role_assign_tbl ura ON r.col_name = ura.col_role_name").
		Where("ura.col_user_id = ?", userID).
		Find(&roles).Error
	
	if err != nil {
		return nil, err
	}
	return roles, nil
}
