package comment

import "context"

type CommentRepo interface {
	PostComment(ctx context.Context, input *Comment) error
	GetCommentByUserId(ctx context.Context, userId uint) ([]Comment, error)
}
