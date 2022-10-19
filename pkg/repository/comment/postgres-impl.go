package comment

import (
	"context"
	"errors"

	customError "github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/errors"

	"github.com/syahrilmaulayahya/final-project-go-fga/config/postgres"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/comment"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/photo"
	"gorm.io/gorm"
)

type CommentRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewCommentRepo(pgCln postgres.PostgresClient) comment.CommentRepo {
	return &CommentRepoImpl{pgCln: pgCln}
}

func (c *CommentRepoImpl) PostComment(ctx context.Context, input *comment.Comment) error {
	var resultPhoto photo.Photo
	db := c.pgCln.GetClient()
	searchPhoto := db.Model(&photo.Photo{}).First(&resultPhoto, "id = ?", input.PhotoID)
	if errors.Is(searchPhoto.Error, gorm.ErrRecordNotFound) {
		return customError.ErrPhotoNotFound
	}
	result := db.Model(&comment.Comment{}).Create(&input)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *CommentRepoImpl) GetCommentByUserId(ctx context.Context, userId uint) ([]comment.Comment, error) {
	var result []comment.Comment
	db := c.pgCln.GetClient()
	resultDb := db.Preload("Photo").Find(&result, "user_id = ?", userId)
	if resultDb.Error != nil {
		return nil, resultDb.Error
	}
	return result, nil
}
