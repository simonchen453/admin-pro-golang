package constants

// 通用状态常量
const (
	// 通用状态
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusLocked   = "locked"
	StatusDeleted  = "deleted"

	// 用户状态
	UserStatusActive   = "active"
	UserStatusInactive = "inactive"
	UserStatusLocked   = "locked"

	// 部门状态
	DeptStatusActive = "active"

	// 字典类型
	DictTypeSystem   = "system"
	DictTypeBusiness = "business"

	// 是否
	Yes = "1"
	No  = "0"
)

// 错误码
const (
	ErrorCodeSuccess      = "200"
	ErrorCodeBadRequest   = "400"
	ErrorCodeUnauthorized = "401"
	ErrorCodeForbidden    = "403"
	ErrorCodeNotFound     = "404"
	ErrorCodeServerError  = "500"
)
