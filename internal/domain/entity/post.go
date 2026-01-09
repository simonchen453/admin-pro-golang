package entity

import "time"

// Post 对应 sys_post_tbl
type Post struct {
	ID          string    `json:"id" gorm:"column:col_id;primaryKey"`
	Code        string    `json:"code" gorm:"column:col_code"`
	Name        string    `json:"name" gorm:"column:col_name"`
	Sort        int       `json:"sort" gorm:"column:col_sort"`
	Status      string    `json:"status" gorm:"column:col_status"`
	Remark      string    `json:"remark" gorm:"column:col_remark"`
	CreatedBy   string    `json:"createdBy" gorm:"column:col_created_by_user_id"`
	CreatedDate time.Time `json:"createdDate" gorm:"column:col_created_date"`
	UpdatedBy   string    `json:"updatedBy" gorm:"column:col_updated_by_user_id"`
	UpdatedDate time.Time `json:"updatedDate" gorm:"column:col_updated_date"`
}

func (Post) TableName() string {
	return "sys_post_tbl"
}
