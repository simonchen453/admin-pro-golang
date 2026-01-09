package entity

// Menu 对应 sys_menu_tbl
type Menu struct {
	ID          string    `json:"id" gorm:"column:col_id;primaryKey"`
	Name        string    `json:"name" gorm:"column:col_name"`
	Display     string    `json:"display" gorm:"column:col_display"`
	ParentID    string    `json:"parentId" gorm:"column:col_parent_id"`
	OrderNum    int       `json:"orderNum" gorm:"column:col_order_num"`
	URL         string    `json:"url" gorm:"column:col_url"`
	Type        string    `json:"type" gorm:"column:col_type"`    // M目录 C菜单 F按钮
	Visible     string    `json:"visible" gorm:"column:col_visible"`
	Status      string    `json:"status" gorm:"column:col_status"`
	Permission  string    `json:"permission" gorm:"column:col_permission"`
	Icon        string    `json:"icon" gorm:"column:col_icon"`
	Children    []*Menu   `json:"children,omitempty" gorm:"-"`
}

func (Menu) TableName() string {
	return "sys_menu_tbl"
}
