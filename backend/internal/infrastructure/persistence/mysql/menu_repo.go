package mysql

import (
	"context"

	"gorm.io/gorm"
	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) repository.MenuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) GetMenusByRoleNames(ctx context.Context, roleNames []string) ([]*entity.Menu, error) {
	var menus []*entity.Menu
	if len(roleNames) == 0 {
		return menus, nil
	}

	// Join sys_role_menu_assign_tbl and sys_menu_tbl on menu_name
	// Distinct to avoid duplicates if multiple roles have same menu
	err := r.db.WithContext(ctx).
		Table("sys_menu_tbl m").
		Distinct("m.*").
		Joins("JOIN sys_role_menu_assign_tbl rma ON m.col_name = rma.col_menu_name").
		Where("rma.col_role_name IN ?", roleNames).
		Order("m.col_order_num ASC").
		Find(&menus).Error
	
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) GetAllMenus(ctx context.Context) ([]*entity.Menu, error) {
	var menus []*entity.Menu
	err := r.db.WithContext(ctx).
		Order("col_order_num ASC").
		Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}
