package mysql

import (
	"context"
	"errors"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) GetList(ctx context.Context) ([]*entity.Post, error) {
	var posts []*entity.Post
	err := r.db.WithContext(ctx).
		Order("col_post_sort ASC").
		Find(&posts).Error
	return posts, err
}

func (r *postRepository) GetByID(ctx context.Context, id string) (*entity.Post, error) {
	var post entity.Post
	err := r.db.WithContext(ctx).Where("col_id = ?", id).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) Create(ctx context.Context, post *entity.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

func (r *postRepository) Update(ctx context.Context, post *entity.Post) error {
	return r.db.WithContext(ctx).Save(post).Error
}

func (r *postRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("col_id = ?", id).Delete(&entity.Post{}).Error
}

func (r *postRepository) CheckPostCodeUnique(ctx context.Context, code string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Post{}).
		Where("col_post_code = ?", code).
		Count(&count).Error
	return count == 0, err
}

func (r *postRepository) CheckPostNameUnique(ctx context.Context, name string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Post{}).
		Where("col_post_name = ?", name).
		Count(&count).Error
	return count == 0, err
}
