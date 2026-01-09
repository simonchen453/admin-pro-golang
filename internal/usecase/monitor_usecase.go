package usecase

import (
	"context"
	"runtime"
	"time"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type MonitorUsecase interface {
	GetServerInfo(ctx context.Context) (*entity.Server, error)
	GetSessionList(ctx context.Context) ([]*entity.Session, error)
	DeleteSession(ctx context.Context, id string) error
	CreateSession(ctx context.Context, session *entity.Session) error
}

type monitorUsecase struct {
	monitorRepo repository.MonitorRepository
}

func NewMonitorUsecase(monitorRepo repository.MonitorRepository) MonitorUsecase {
	return &monitorUsecase{
		monitorRepo: monitorRepo,
	}
}

func (u *monitorUsecase) GetServerInfo(ctx context.Context) (*entity.Server, error) {
	server := &entity.Server{}

	// CPU
	cpuPercent, _ := cpu.Percent(0, false)
	if len(cpuPercent) > 0 {
		server.CPU.Used = cpuPercent[0]
	}
	server.CPU.Cores, _ = cpu.Counts(true)

	// Memory
	vMem, _ := mem.VirtualMemory()
	server.Mem.Total = vMem.Total
	server.Mem.Used = vMem.Used
	server.Mem.Free = vMem.Free
	server.Mem.Usage = vMem.UsedPercent

	// Disk
	parts, _ := disk.Partitions(true)
	diskTotal := uint64(0)
	diskUsed := uint64(0)
	for _, part := range parts {
		usage, _ := disk.Usage(part.Mountpoint)
		if usage != nil {
			diskTotal += usage.Total
			diskUsed += usage.Used
		}
	}
	server.Disk.Total = diskTotal
	server.Disk.Used = diskUsed
	server.Disk.Free = diskTotal - diskUsed
	if diskTotal > 0 {
		server.Disk.Usage = float64(diskUsed) / float64(diskTotal) * 100
	}

	// Go
	server.Go.Version = runtime.Version()
	server.Go.NumGoroutine = runtime.NumGoroutine()

	return server, nil
}

func (u *monitorUsecase) GetSessionList(ctx context.Context) ([]*entity.Session, error) {
	return u.monitorRepo.GetSessionList(ctx)
}

func (u *monitorUsecase) DeleteSession(ctx context.Context, id string) error {
	return u.monitorRepo.DeleteSession(ctx, id)
}

func (u *monitorUsecase) CreateSession(ctx context.Context, session *entity.Session) error {
	session.CreatedDate = time.Now()
	return u.monitorRepo.CreateSession(ctx, session)
}
