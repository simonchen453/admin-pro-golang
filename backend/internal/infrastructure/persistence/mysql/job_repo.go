package mysql

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) repository.JobRepository {
	return &jobRepository{db: db}
}

func (r *jobRepository) GetJobList(ctx context.Context) ([]*entity.Job, error) {
	var list []*entity.Job
	err := r.db.WithContext(ctx).Find(&list).Error
	return list, err
}

func (r *jobRepository) GetJob(ctx context.Context, id string) (*entity.Job, error) {
	var job entity.Job
	err := r.db.WithContext(ctx).Where("col_id = ?", id).First(&job).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &job, err
}

func (r *jobRepository) CreateJob(ctx context.Context, job *entity.Job) error {
	return r.db.WithContext(ctx).Create(job).Error
}

func (r *jobRepository) UpdateJob(ctx context.Context, job *entity.Job) error {
	return r.db.WithContext(ctx).Save(job).Error
}

func (r *jobRepository) DeleteJob(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("col_id = ?", id).Delete(&entity.Job{}).Error
}

func (r *jobRepository) GetJobLogList(ctx context.Context) ([]*entity.JobLog, error) {
	var list []*entity.JobLog
	err := r.db.WithContext(ctx).Order("col_created_time DESC").Limit(100).Find(&list).Error
	return list, err
}

func (r *jobRepository) GetJobLog(ctx context.Context, id string) (*entity.JobLog, error) {
	var log entity.JobLog
	err := r.db.WithContext(ctx).Where("col_id = ?", id).First(&log).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &log, err
}
