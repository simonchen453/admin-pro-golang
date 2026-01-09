package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

type JobUsecase interface {
	GetJobList(ctx context.Context) ([]*entity.Job, error)
	GetJob(ctx context.Context, id string) (*entity.Job, error)
	CreateJob(ctx context.Context, job *entity.Job) error
	UpdateJob(ctx context.Context, job *entity.Job) error
	DeleteJob(ctx context.Context, id string) error
	
	GetJobLogList(ctx context.Context) ([]*entity.JobLog, error)
	GetJobLog(ctx context.Context, id string) (*entity.JobLog, error)
}

type jobUsecase struct {
	jobRepo repository.JobRepository
}

func NewJobUsecase(jobRepo repository.JobRepository) JobUsecase {
	return &jobUsecase{
		jobRepo: jobRepo,
	}
}

func (u *jobUsecase) GetJobList(ctx context.Context) ([]*entity.Job, error) {
	return u.jobRepo.GetJobList(ctx)
}

func (u *jobUsecase) GetJob(ctx context.Context, id string) (*entity.Job, error) {
	return u.jobRepo.GetJob(ctx, id)
}

func (u *jobUsecase) CreateJob(ctx context.Context, job *entity.Job) error {
	job.ID = uuid.NewString()
	job.CreatedTime = time.Now()
	// TODO: Add to scheduler
	return u.jobRepo.CreateJob(ctx, job)
}

func (u *jobUsecase) UpdateJob(ctx context.Context, job *entity.Job) error {
	exist, err := u.jobRepo.GetJob(ctx, job.ID)
	if err != nil {
		return err
	}
	if exist == nil {
		return errors.New("job not found")
	}
	// TODO: Update scheduler
	return u.jobRepo.UpdateJob(ctx, job)
}

func (u *jobUsecase) DeleteJob(ctx context.Context, id string) error {
	// TODO: Remove from scheduler
	return u.jobRepo.DeleteJob(ctx, id)
}

func (u *jobUsecase) GetJobLogList(ctx context.Context) ([]*entity.JobLog, error) {
	return u.jobRepo.GetJobLogList(ctx)
}

func (u *jobUsecase) GetJobLog(ctx context.Context, id string) (*entity.JobLog, error) {
	return u.jobRepo.GetJobLog(ctx, id)
}
