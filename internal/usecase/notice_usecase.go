package usecase

import (
	"context"
	"errors"
	"time"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
	"github.com/google/uuid"
)

type NoticeUsecase interface {
	GetNoticeList(ctx context.Context) ([]*entity.Notice, error)
	GetNotice(ctx context.Context, id string) (*entity.Notice, error)
	CreateNotice(ctx context.Context, notice *entity.Notice) error
	UpdateNotice(ctx context.Context, notice *entity.Notice) error
	DeleteNotice(ctx context.Context, id string) error
}

type noticeUsecase struct {
	noticeRepo repository.NoticeRepository
}

func NewNoticeUsecase(noticeRepo repository.NoticeRepository) NoticeUsecase {
	return &noticeUsecase{
		noticeRepo: noticeRepo,
	}
}

func (u *noticeUsecase) GetNoticeList(ctx context.Context) ([]*entity.Notice, error) {
	return u.noticeRepo.GetList(ctx)
}

func (u *noticeUsecase) GetNotice(ctx context.Context, id string) (*entity.Notice, error) {
	return u.noticeRepo.GetByID(ctx, id)
}

func (u *noticeUsecase) CreateNotice(ctx context.Context, notice *entity.Notice) error {
	notice.ID = uuid.NewString()
	notice.CreatedDate = time.Now()
	// Default logic if needed
	return u.noticeRepo.Create(ctx, notice)
}

func (u *noticeUsecase) UpdateNotice(ctx context.Context, notice *entity.Notice) error {
	exist, err := u.noticeRepo.GetByID(ctx, notice.ID)
	if err != nil {
		return err
	}
	if exist == nil {
		return errors.New("notice not found")
	}
	notice.UpdatedDate = time.Now()
	return u.noticeRepo.Update(ctx, notice)
}

func (u *noticeUsecase) DeleteNotice(ctx context.Context, id string) error {
	return u.noticeRepo.Delete(ctx, id)
}
