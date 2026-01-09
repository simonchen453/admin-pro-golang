package usecase

import (
	"context"
	"errors"
	"time"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
	"admin-pro/pkg/scheduler"
	"github.com/google/uuid"
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
	jobRepo   repository.JobRepository
	scheduler *scheduler.Scheduler
}

func NewJobUsecase(jobRepo repository.JobRepository, sched *scheduler.Scheduler) JobUsecase {
	return &jobUsecase{
		jobRepo:   jobRepo,
		scheduler: sched,
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

	// Create job execution function
	jobFunc := func(ctx context.Context) error {
		// Job execution logic - can be extended based on job type
		return u.executeJob(ctx, job)
	}

	// Add to scheduler
	if err := u.scheduler.AddJob(job, jobFunc); err != nil {
		return err
	}

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

	// Create job execution function
	jobFunc := func(ctx context.Context) error {
		return u.executeJob(ctx, job)
	}

	// Update scheduler
	if err := u.scheduler.UpdateJob(job, jobFunc); err != nil {
		return err
	}

	return u.jobRepo.UpdateJob(ctx, job)
}

func (u *jobUsecase) DeleteJob(ctx context.Context, id string) error {
	// Remove from scheduler
	u.scheduler.RemoveJob(id)
	return u.jobRepo.DeleteJob(ctx, id)
}

func (u *jobUsecase) GetJobLogList(ctx context.Context) ([]*entity.JobLog, error) {
	return u.jobRepo.GetJobLogList(ctx)
}

func (u *jobUsecase) GetJobLog(ctx context.Context, id string) (*entity.JobLog, error) {
	return u.jobRepo.GetJobLog(ctx, id)
}

// executeJob executes the job logic
func (u *jobUsecase) executeJob(ctx context.Context, job *entity.Job) error {
	// Execute the job (placeholder - actual implementation depends on job type)
	// This is where you would call the actual job execution logic
	// For now, we just simulate execution
	_ = ctx // TODO: Use context for cancellation

	// TODO: Implement actual job execution based on BeanName and MethodName
	// For example: use reflection or a registry pattern to call the actual method
	// This is a placeholder that simulates successful job execution

	return nil
}
