package entity

// Menu 对应 sys_menu_tbl
type Menu struct {
	ID         string  `json:"id" gorm:"column:col_id;primaryKey" validate:"omitempty,uuid"`
	Name       string  `json:"name" gorm:"column:col_name" validate:"required,min=2,max=50"`
	Display    string  `json:"display" gorm:"column:col_display" validate:"required,min=2,max=100"`
	ParentID   string  `json:"parentId" gorm:"column:col_parent_id" validate:"omitempty,uuid"`
	OrderNum   int     `json:"orderNum" gorm:"column:col_order_num" validate:"gte=0"`
	URL        string  `json:"url" gorm:"column:col_url" validate:"omitempty,max=255"`
	Type       string  `json:"type" gorm:"column:col_type" validate:"required,oneof=M C F"` // M目录 C菜单 F按钮
	Visible    string  `json:"visible" gorm:"column:col_visible" validate:"omitempty,oneof=0 1"`
	Status     string  `json:"status" gorm:"column:col_status" validate:"omitempty,oneof=0 1"`
	Permission string  `json:"permission" gorm:"column:col_permission" validate:"omitempty,max=100"`
	Icon       string  `json:"icon" gorm:"column:col_icon" validate:"omitempty,max=50"`
	Children   []*Menu `json:"children,omitempty" gorm:"-"`
}

func (Menu) TableName() string {
	return "sys_menu_tbl"
}
