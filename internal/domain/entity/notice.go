package entity

import "time"

// Notice 对应 sys_notification_tbl
type Notice struct {
	ID          string     `json:"id" gorm:"column:col_id;primaryKey"`
	Title       string     `json:"title" gorm:"column:col_title"`
	Content     string     `json:"content" gorm:"column:col_content"`
	UserDomain  string     `json:"userDomain" gorm:"column:col_user_domain"`
	StartTime   *time.Time `json:"startTime" gorm:"column:col_start_time"`
	EndTime     *time.Time `json:"endTime" gorm:"column:col_end_time"`
	CreatedBy   string     `json:"createdBy" gorm:"column:col_created_by_user_id"`
	CreatedDate time.Time  `json:"createdDate" gorm:"column:col_created_date"`
	UpdatedBy   string     `json:"updatedBy" gorm:"column:col_updated_by_user_id"`
	UpdatedDate time.Time  `json:"updatedDate" gorm:"column:col_updated_date"`
}

func (Notice) TableName() string {
	return "sys_notification_tbl"
}
