package usecase

import (
	"context"

	"admin-pro/internal/config"
	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
	apperrors "admin-pro/pkg/errors"
	"admin-pro/pkg/utils"
)

type AuthUsecase interface {
	Login(ctx context.Context, loginName, password string) (string, *entity.User, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
	cfg      *config.Config
}

func NewAuthUsecase(userRepo repository.UserRepository, cfg *config.Config) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (u *authUsecase) Login(ctx context.Context, loginName, password string) (string, *entity.User, error) {
	user, err := u.userRepo.GetByLoginName(ctx, loginName)
	if err != nil {
		return "", nil, apperrors.Wrap(err, "failed to get user by login name")
	}
	if user == nil {
		return "", nil, apperrors.ErrUserNotFound
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", nil, apperrors.ErrInvalidPassword
	}

	token, err := utils.GenerateToken(user.UserID, user.UserDomain, user.LoginName, u.cfg)
	if err != nil {
		return "", nil, apperrors.Wrap(err, "failed to generate token")
	}

	return token, user, nil
}
