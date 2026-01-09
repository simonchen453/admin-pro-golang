package repository

import (
	"admin-pro/internal/domain/entity"
	"context"
)

type MonitorRepository interface {
	GetSessionList(ctx context.Context) ([]*entity.Session, error)
	DeleteSession(ctx context.Context, id string) error
	CreateSession(ctx context.Context, session *entity.Session) error
}
