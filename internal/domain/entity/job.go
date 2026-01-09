package entity

import "time"

// Job 对应 sys_schedule_job_tbl
type Job struct {
	ID             string    `json:"id" gorm:"column:col_id;primaryKey" validate:"omitempty,uuid"`
	BeanName       string    `json:"beanName" gorm:"column:col_bean_name" validate:"required,min=2,max=100"`
	MethodName     string    `json:"methodName" gorm:"column:col_method_name" validate:"required,min=2,max=100"`
	Params         string    `json:"params" gorm:"column:col_params" validate:"omitempty,max=500"`
	CronExpression string    `json:"cronExpression" gorm:"column:col_cron_expression" validate:"required,max=100"`
	Status         int       `json:"status" gorm:"column:col_status" validate:"oneof=0 1"` // 0:正常 1:暂停
	Remark         string    `json:"remark" gorm:"column:col_remark" validate:"omitempty,max=500"`
	CreatedTime    time.Time `json:"createdTime" gorm:"column:col_created_time"`
}

func (Job) TableName() string {
	return "sys_schedule_job_tbl"
}

// JobLog 对应 sys_schedule_job_log_tbl
type JobLog struct {
	ID          string    `json:"id" gorm:"column:col_id;primaryKey"`
	JobID       string    `json:"jobId" gorm:"column:col_job_id"`
	BeanName    string    `json:"beanName" gorm:"column:col_bean_name"`
	MethodName  string    `json:"methodName" gorm:"column:col_method_name"`
	Params      string    `json:"params" gorm:"column:col_params"`
	Status      int       `json:"status" gorm:"column:col_status"` // 0:成功 1:失败
	Error       string    `json:"error" gorm:"column:col_error"`
	Times       int       `json:"times" gorm:"column:col_times"` // 耗时(Int)
	CreatedTime time.Time `json:"createdTime" gorm:"column:col_created_time"`
}

func (JobLog) TableName() string {
	return "sys_schedule_job_log_tbl"
}
