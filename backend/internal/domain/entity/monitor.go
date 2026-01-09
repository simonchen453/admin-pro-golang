package entity

import "time"

// Session 对应 sys_session_tbl
type Session struct {
	ID                  string    `json:"id" gorm:"column:col_id;primaryKey"`
	SessionID           string    `json:"sessionId" gorm:"column:col_session_id"`
	ThirdPartySessionID string    `json:"thirdPartySessionId" gorm:"column:col_third_party_session_id"`
	Status              string    `json:"status" gorm:"column:col_status"`
	UserDomain          string    `json:"userDomain" gorm:"column:col_user_domain"`
	UserID              string    `json:"userId" gorm:"column:col_user_id"`
	LoginName           string    `json:"loginName" gorm:"column:col_login_name"`
	IPAddr              string    `json:"ipAddr" gorm:"column:col_ip_addr"`
	LoginLocation       string    `json:"loginLocation" gorm:"column:col_login_location"`
	Browser             string    `json:"browser" gorm:"column:col_browser"`
	OS                  string    `json:"os" gorm:"column:col_os"`
	DeptNo              string    `json:"deptNo" gorm:"column:col_dept_no"`

	CreatedBy   string    `json:"createdBy" gorm:"column:col_created_by_user_id"`
	CreatedDate time.Time `json:"createdDate" gorm:"column:col_created_date"`
}

func (Session) TableName() string {
	return "sys_session_tbl"
}

// Server info struct (not DB table)
type Server struct {
	CPU  CPUInfo  `json:"cpu"`
	Mem  MemInfo  `json:"mem"`
	Disk DiskInfo `json:"disk"`
	Go   GoInfo   `json:"go"`
}

type CPUInfo struct {
	Cores int     `json:"cores"`
	Used  float64 `json:"used"`
}

type MemInfo struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
	Usage float64 `json:"usage"`
}

type DiskInfo struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
	Usage float64 `json:"usage"`
}

type GoInfo struct {
	Version string `json:"version"`
	NumGoroutine int `json:"numGoroutine"`
}
