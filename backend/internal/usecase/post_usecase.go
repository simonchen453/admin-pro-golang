package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

type PostUsecase interface {
	GetPostList(ctx context.Context) ([]*entity.Post, error)
	GetPost(ctx context.Context, id string) (*entity.Post, error)
	CreatePost(ctx context.Context, post *entity.Post) error
	UpdatePost(ctx context.Context, post *entity.Post) error
	DeletePost(ctx context.Context, id string) error
}

type postUsecase struct {
	postRepo repository.PostRepository
}

func NewPostUsecase(postRepo repository.PostRepository) PostUsecase {
	return &postUsecase{
		postRepo: postRepo,
	}
}

func (u *postUsecase) GetPostList(ctx context.Context) ([]*entity.Post, error) {
	return u.postRepo.GetList(ctx)
}

func (u *postUsecase) GetPost(ctx context.Context, id string) (*entity.Post, error) {
	return u.postRepo.GetByID(ctx, id)
}

func (u *postUsecase) CreatePost(ctx context.Context, post *entity.Post) error {
	isUnique, err := u.postRepo.CheckPostCodeUnique(ctx, post.Code)
	if err != nil {
		return err
	}
	if !isUnique {
		return errors.New("post code already exists")
	}
	
	post.ID = uuid.NewString()
	post.CreatedDate = time.Now()
	post.Status = "active"
	return u.postRepo.Create(ctx, post)
}

func (u *postUsecase) UpdatePost(ctx context.Context, post *entity.Post) error {
	existing, err := u.postRepo.GetByID(ctx, post.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("post not found")
	}

	post.UpdatedDate = time.Now()
	return u.postRepo.Update(ctx, post)
}

func (u *postUsecase) DeletePost(ctx context.Context, id string) error {
	return u.postRepo.Delete(ctx, id)
}
