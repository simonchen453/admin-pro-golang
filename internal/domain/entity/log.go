package entity

import "time"

// LoginLog 对应 sys_login_info_tbl
type LoginLog struct {
	ID            string    `json:"id" gorm:"column:col_id;primaryKey"`
	UserID        string    `json:"userId" gorm:"column:col_user_id"`
	UserDomain    string    `json:"userDomain" gorm:"column:col_user_domain"`
	Key           string    `json:"key" gorm:"column:col_key"` // 例如： "admin"
	IPAddress     string    `json:"ipAddress" gorm:"column:col_ip_address"`
	LoginLocation string    `json:"loginLocation" gorm:"column:col_login_location"`
	Browser       string    `json:"browser" gorm:"column:col_browser"`
	OS            string    `json:"os" gorm:"column:col_os"`
	Status        string    `json:"status" gorm:"column:col_status"`
	Message       string    `json:"message" gorm:"column:col_message"`
	LoginTime     time.Time `json:"loginTime" gorm:"column:col_login_time"`

	CreatedBy   string    `json:"createdBy" gorm:"column:col_created_by_user_id"`
	CreatedDate time.Time `json:"createdDate" gorm:"column:col_created_date"`
}

func (LoginLog) TableName() string {
	return "sys_login_info_tbl"
}

// OperLog 对应 sys_audit_log_tbl
type OperLog struct {
	ID        string    `json:"id" gorm:"column:col_id;primaryKey"`
	LogDate   time.Time `json:"logDate" gorm:"column:col_log_date"`
	Category  string    `json:"category" gorm:"column:col_category"`
	Module    string    `json:"module" gorm:"column:col_module"`
	IPAddress string    `json:"ipAddress" gorm:"column:col_ip_address"`
	Status    string    `json:"status" gorm:"column:col_status"`
	Event     string    `json:"event" gorm:"column:col_event"`
	EventData string    `json:"eventData" gorm:"column:col_event_data"`
	SessionID string    `json:"sessionId" gorm:"column:col_session_id"`

	CreatedBy   string    `json:"createdBy" gorm:"column:col_created_by_user_id"`
	CreatedDate time.Time `json:"createdDate" gorm:"column:col_created_date"`
}

func (OperLog) TableName() string {
	return "sys_audit_log_tbl"
}
