package comment

import "github.com/gin-gonic/gin"

type CommentHandler interface {
	PostCommentHdl(ctx *gin.Context)
	GetCommentByUserIdHdl(ctx *gin.Context)
	EditCommentHdl(ctx *gin.Context)
	DeleteCommentHdl(ctx *gin.Context)
}
