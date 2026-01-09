package entity

import "time"

// Role 对应 sys_role_tbl
type Role struct {
	ID          string    `json:"id" gorm:"column:col_id;primaryKey"`
	Name        string    `json:"name" gorm:"column:col_name"`
	Display     string    `json:"display" gorm:"column:col_display"`
	Status      string    `json:"status" gorm:"column:col_status"`
	IsSystem    *bool     `json:"isSystem" gorm:"column:col_is_system"`
	CreatedDate time.Time `json:"createdDate" gorm:"column:col_created_date"`
}

func (Role) TableName() string {
	return "sys_role_tbl"
}
