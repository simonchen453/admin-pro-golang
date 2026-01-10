package errors

import (
	"fmt"
)

// 通用错误定义
var (
	// 认证错误
	ErrUserNotFound    = fmt.Errorf("user not found")
	ErrInvalidPassword = fmt.Errorf("invalid password")
	ErrInvalidToken    = fmt.Errorf("invalid token")
	ErrTokenExpired    = fmt.Errorf("token expired")
	ErrUnauthorized    = fmt.Errorf("unauthorized access")

	// 校验错误
	ErrInvalidInput  = fmt.Errorf("invalid input")
	ErrInvalidParams = fmt.Errorf("invalid parameters")
	ErrRequiredField = fmt.Errorf("required field missing")

	// 资源错误
	ErrNotFound      = fmt.Errorf("resource not found")
	ErrAlreadyExists = fmt.Errorf("resource already exists")
	ErrConflict      = fmt.Errorf("resource conflict")

	// 业务逻辑错误
	ErrInvalidOperation = fmt.Errorf("invalid operation")
	ErrPermissionDenied = fmt.Errorf("permission denied")
	ErrStatusInvalid    = fmt.Errorf("invalid status")

	// 数据库错误
	ErrDatabase       = fmt.Errorf("database error")
	ErrRecordNotFound = fmt.Errorf("record not found")
	ErrDuplicateEntry = fmt.Errorf("duplicate entry")
)

// Wrap 包装错误，添加额外上下文
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// Wrapf 包装错误，使用格式化消息
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
}
