package constants

// 状态常量
const (
	StatusActive   = "active"
	StatusDisabled = "disabled"
	StatusDeleted  = "deleted"
)

// 用户状态
const (
	UserStatusActived  = "Actived"
	UserStatusDisabled = "Disabled"
)

// 菜单类型
const (
	MenuTypeDirectory = "M" // 目录
	MenuTypeMenu      = "C" // 菜单
	MenuTypeButton    = "F" // 按钮
)

// 可见性
const (
	VisibleShow = "0"
	VisibleHide = "1"
)

// 任务状态
const (
	JobStatusNormal = 0 // 正常
	JobStatusPaused = 1 // 暂停
)

// 任务执行状态
const (
	JobLogStatusSuccess = 0 // 成功
	JobLogStatusFailed  = 1 // 失败
)

// 是否系统内置
const (
	IsSystemYes = true
	IsSystemNo  = false
)
