package entity

import "time"

// User 实体，对应数据库中的 sys_user_tbl 表
// 结构体中的字段通过 gorm tag 与数据库列一一对应
type User struct {
	ID                 string    `json:"id" gorm:"column:col_id;primaryKey"`           // 主键 ID
	UserDomain         string    `json:"userDomain" gorm:"column:col_user_domain"`     // 用户域 (租户隔离用)
	UserID             string    `json:"userId" gorm:"column:col_user_id"`             // 用户 ID (业务主键)
	LoginName          string    `json:"loginName" gorm:"column:col_login_name"`       // 登录名
	Display            string    `json:"display" gorm:"column:col_display"`            // 显示名称
	MobileNo           string    `json:"mobileNo" gorm:"column:col_mobile_no"`         // 手机号
	Password           string    `json:"-" gorm:"column:col_pwd"`                      // 密码 (json:"-" 表示序列化时不返回给前端)
	Email              string    `json:"email" gorm:"column:col_email"`                // 邮箱
	Status             string    `json:"status" gorm:"column:col_status"`              // 状态 (Actived/Disabled)
	AvatarUrl          string    `json:"avatarUrl" gorm:"column:col_avatar_url"`       // 头像地址
	IsSystem           *bool     `json:"isSystem" gorm:"column:col_is_system"`         // 是否系统用户 (不可删除)
	CreatedByUserDomain string   `json:"createdByUserDomain" gorm:"column:col_created_by_user_domain"`
	CreatedByUserID     string   `json:"createdByUserId" gorm:"column:col_created_by_user_id"`
	CreatedDate         time.Time `json:"createdDate" gorm:"column:col_created_date"` // 创建时间
}

// TableName 指定 User 结构体对应的数据库表名
func (User) TableName() string {
	return "sys_user_tbl"
}
