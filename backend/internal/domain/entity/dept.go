package entity

import "time"

// Dept 对应 sys_dept_tbl
type Dept struct {
	ID          string    `json:"id" gorm:"column:col_id;primaryKey"`
	Name        string    `json:"name" gorm:"column:col_name"`
	ParentID    string    `json:"parentId" gorm:"column:col_parent_id"`
	OrderNum    int       `json:"orderNum" gorm:"column:col_order_num"`
	Leader      string    `json:"leader" gorm:"column:col_leader"`
	Phone       string    `json:"phone" gorm:"column:col_phone"`
	Email       string    `json:"email" gorm:"column:col_email"`
	Status      string    `json:"status" gorm:"column:col_status"`
	CreatedBy   string    `json:"createdBy" gorm:"column:col_created_by_user_id"`
	CreatedDate time.Time `json:"createdDate" gorm:"column:col_created_date"`
	UpdatedBy   string    `json:"updatedBy" gorm:"column:col_updated_by_user_id"`
	UpdatedDate time.Time `json:"updatedDate" gorm:"column:col_updated_date"`
	Children    []*Dept   `json:"children,omitempty" gorm:"-"`
}

func (Dept) TableName() string {
	return "sys_dept_tbl"
}
