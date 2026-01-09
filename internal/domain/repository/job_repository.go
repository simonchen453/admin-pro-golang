package repository

import (
	"admin-pro/internal/domain/entity"
	"context"
)

type JobRepository interface {
	GetJobList(ctx context.Context) ([]*entity.Job, error)
	GetJob(ctx context.Context, id string) (*entity.Job, error)
	CreateJob(ctx context.Context, job *entity.Job) error
	UpdateJob(ctx context.Context, job *entity.Job) error
	DeleteJob(ctx context.Context, id string) error

	// Log
	GetJobLogList(ctx context.Context) ([]*entity.JobLog, error)
	GetJobLog(ctx context.Context, id string) (*entity.JobLog, error)
}
