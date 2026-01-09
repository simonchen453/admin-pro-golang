package usecase

import (
	"context"
	"errors"

	"admin-pro/internal/config"
	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
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
		return "", nil, err
	}
	if user == nil {
		return "", nil, errors.New("user not found")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", nil, errors.New("invalid password")
	}

	token, err := utils.GenerateToken(user.UserID, user.UserDomain, user.LoginName, u.cfg)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
