package scheduler

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"admin-pro/internal/domain/entity"
	"github.com/robfig/cron/v3"
)

// JobFunc 定时任务执行函数类型
type JobFunc func(ctx context.Context) error

// Scheduler 定时任务调度器
type Scheduler struct {
	cron     *cron.Cron
	jobs     map[string]cron.EntryID
	jobFuncs map[string]JobFunc
	mu       sync.RWMutex
}

// NewScheduler 创建新的调度器
func NewScheduler() *Scheduler {
	return &Scheduler{
		cron:     cron.New(cron.WithSeconds()), // 支持秒级 cron
		jobs:     make(map[string]cron.EntryID),
		jobFuncs: make(map[string]JobFunc),
	}
}

// Start 启动调度器
func (s *Scheduler) Start() {
	s.cron.Start()
	log.Println("Scheduler started")
}

// Stop 停止调度器
func (s *Scheduler) Stop() {
	s.cron.Stop()
	log.Println("Scheduler stopped")
}

// AddJob 添加定时任务
func (s *Scheduler) AddJob(job *entity.Job, jobFunc JobFunc) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	jobID := job.ID

	// 移除旧任务（如果存在）
	if entryID, exists := s.jobs[jobID]; exists {
		s.cron.Remove(entryID)
	}

	// 添加新任务
	entryID, err := s.cron.AddFunc(job.CronExpression, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		jobName := fmt.Sprintf("%s.%s", job.BeanName, job.MethodName)
		log.Printf("Executing job: %s", jobName)
		if err := jobFunc(ctx); err != nil {
			log.Printf("Job %s execution failed: %v", jobName, err)
		} else {
			log.Printf("Job %s executed successfully", jobName)
		}
	})

	if err != nil {
		return fmt.Errorf("failed to add job: %w", err)
	}

	s.jobs[jobID] = entryID
	s.jobFuncs[jobID] = jobFunc
	log.Printf("Job added: %s with cron: %s", jobID, job.CronExpression)
	return nil
}

// RemoveJob 移除定时任务
func (s *Scheduler) RemoveJob(jobID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if entryID, exists := s.jobs[jobID]; exists {
		s.cron.Remove(entryID)
		delete(s.jobs, jobID)
		delete(s.jobFuncs, jobID)
		log.Printf("Job removed: %s", jobID)
	}
}

// UpdateJob 更新定时任务
func (s *Scheduler) UpdateJob(job *entity.Job, jobFunc JobFunc) error {
	return s.AddJob(job, jobFunc)
}

// GetJobCount 获取任务数量
func (s *Scheduler) GetJobCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.jobs)
}
