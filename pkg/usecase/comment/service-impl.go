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

func (c *CommentUseCaseImplt) GetCommentByUserIdSvc(ctx context.Context, userId uint) ([]comment.Comment, error) {
	var result []comment.Comment
	var err error

	result, err = c.commentRepo.GetCommentByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CommentUseCaseImplt) EditCommentSvc(ctx context.Context, input comment.Comment) (comment.Comment, error) {
	var result comment.Comment
	var err error
	if input.Message == "" {
		return comment.Comment{}, errors.ErrMessageEmpty
	}
	result, err = c.commentRepo.EditComment(ctx, input)
	if err != nil {
		return comment.Comment{}, err
	}
	return result, nil
}
