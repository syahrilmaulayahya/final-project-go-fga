package comment

import (
	"context"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/comment"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/errors"
)

type CommentUseCaseImplt struct {
	commentRepo comment.CommentRepo
}

func NewCommentUsecase(commentRepo comment.CommentRepo) comment.CommentUsecase {
	return &CommentUseCaseImplt{commentRepo: commentRepo}
}

func (c *CommentUseCaseImplt) PostCommentSvc(ctx context.Context, input comment.Comment) (comment.Comment, error) {
	if input.Message == "" {
		return comment.Comment{}, errors.ErrMessageEmpty
	}
	err := c.commentRepo.PostComment(ctx, &input)
	if err != nil {
		return comment.Comment{}, err
	}
	return input, nil
}
