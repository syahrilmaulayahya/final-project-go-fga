package comment

import "context"

type CommentUsecase interface {
	PostCommentSvc(ctx context.Context, input Comment) (Comment, error)
}
