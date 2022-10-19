package comment

import (
	"context"

	"github.com/syahrilmaulayahya/final-project-go-fga/config/postgres"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/comment"
)

type CommentRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewCommentRepo(pgCln postgres.PostgresClient) comment.CommentRepo {
	return &CommentRepoImpl{pgCln: pgCln}
}

func (c *CommentRepoImpl) PostComment(ctx context.Context, input *comment.Comment) error {
	db := c.pgCln.GetClient()
	result := db.Model(&comment.Comment{}).Create(&input)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
