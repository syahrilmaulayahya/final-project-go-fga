package comment

import "context"

type CommentUsecase interface {
	PostCommentSvc(ctx context.Context, input Comment) (Comment, error)
	GetCommentByUserIdSvc(ctx context.Context, userId uint) ([]Comment, error)
	EditCommentSvc(ctx context.Context, input Comment) (Comment, error)
}
