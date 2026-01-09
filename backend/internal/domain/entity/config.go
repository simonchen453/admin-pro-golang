package entity

import "time"

// Config 对应 sys_config_tbl
type Config struct {
	ID          string    `json:"id" gorm:"column:col_id;primaryKey"`
	Name        string    `json:"name" gorm:"column:col_name"`
	Key         string    `json:"key" gorm:"column:col_key"`
	Value       string    `json:"value" gorm:"column:col_value"`
	IsSystem    int       `json:"isSystem" gorm:"column:col_system"` // 1=Yes 0=No
	Remark      string    `json:"remark" gorm:"column:col_remark"`
	CreatedBy   string    `json:"createdBy" gorm:"column:col_created_by_user_id"`
	CreatedDate time.Time `json:"createdDate" gorm:"column:col_created_date"`
	UpdatedBy   string    `json:"updatedBy" gorm:"column:col_updated_by_user_id"`
	UpdatedDate time.Time `json:"updatedDate" gorm:"column:col_updated_date"`
}

func (Config) TableName() string {
	return "sys_config_tbl"
}
