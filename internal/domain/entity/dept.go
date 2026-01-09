package entity

import "time"

// Dept 对应 sys_dept_tbl
type Dept struct {
	ID          string    `json:"id" gorm:"column:col_id;primaryKey" validate:"omitempty,uuid"`
	Name        string    `json:"name" gorm:"column:col_name" validate:"required,min=2,max=50"`
	ParentID    string    `json:"parentId" gorm:"column:col_parent_id" validate:"omitempty,uuid"`
	OrderNum    int       `json:"orderNum" gorm:"column:col_order_num" validate:"gte=0"`
	Leader      string    `json:"leader" gorm:"column:col_leader" validate:"omitempty,max=50"`
	Phone       string    `json:"phone" gorm:"column:col_phone" validate:"omitempty,len=11"`
	Email       string    `json:"email" gorm:"column:col_email" validate:"omitempty,email,max=100"`
	Status      string    `json:"status" gorm:"column:col_status" validate:"omitempty,oneof=active inactive"`
	CreatedBy   string    `json:"createdBy" gorm:"column:col_created_by_user_id"`
	CreatedDate time.Time `json:"createdDate" gorm:"column:col_created_date"`
	UpdatedBy   string    `json:"updatedBy" gorm:"column:col_updated_by_user_id"`
	UpdatedDate time.Time `json:"updatedDate" gorm:"column:col_updated_date"`
	Children    []*Dept   `json:"children,omitempty" gorm:"-"`
}

func (Dept) TableName() string {
	return "sys_dept_tbl"
}
