package logger

import (
	"log"
	"os"
)

// LogLevel 代表日志条目的严重程度
type LogLevel int

const (
	// DEBUG 用于详细信息，通常仅在诊断问题时感兴趣
	DEBUG LogLevel = iota
	// INFO 用于一般信息消息
	INFO
	// WARN 用于警告消息
	WARN
	// ERROR 用于错误消息
	ERROR
)

// Logger 是一个简单的结构化记录器
// 在生产环境中，考虑使用更高级的记录器，如 zap 或 zerolog
type Logger struct {
	prefix string
	level  LogLevel
}

// New 创建带有给定前缀的新记录器
func New(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
		level:  INFO,
	}
}

// SetLevel 设置最小日志级别
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Debug 记录调试消息
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.level <= DEBUG {
		log.Printf("[DEBUG] "+l.prefix+" "+format, args...)
	}
}

// Info 记录信息消息
func (l *Logger) Info(format string, args ...interface{}) {
	if l.level <= INFO {
		log.Printf("[INFO] "+l.prefix+" "+format, args...)
	}
}

// Warn 记录警告消息
func (l *Logger) Warn(format string, args ...interface{}) {
	if l.level <= WARN {
		log.Printf("[WARN] "+l.prefix+" "+format, args...)
	}
}

// Error 记录错误消息
func (l *Logger) Error(format string, args ...interface{}) {
	if l.level <= ERROR {
		log.Printf("[ERROR] "+l.prefix+" "+format, args...)
	}
}

// With 返回带有给定前缀的新记录器
func (l *Logger) With(prefix string) *Logger {
	return &Logger{
		prefix: l.prefix + " " + prefix,
		level:  l.level,
	}
}

// 全局记录器实例
var std = New("")

// SetLevel 设置全局记录器的级别
func SetLevel(level LogLevel) {
	std.SetLevel(level)
}

// Debug 使用全局记录器记录调试消息
func Debug(format string, args ...interface{}) {
	std.Debug(format, args...)
}

// Info 使用全局记录器记录信息消息
func Info(format string, args ...interface{}) {
	std.Info(format, args...)
}

// Warn 使用全局记录器记录警告消息
func Warn(format string, args ...interface{}) {
	std.Warn(format, args...)
}

// Error 使用全局记录器记录错误消息
func Error(format string, args ...interface{}) {
	std.Error(format, args...)
}

// With 使用全局记录器返回带有给定前缀的新记录器
func With(prefix string) *Logger {
	return std.With(prefix)
}

// Init 初始化全局记录器
func Init() {
	// 设置输出到标准输出
	log.SetOutput(os.Stdout)
	// 设置标志以包含日期和时间
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
