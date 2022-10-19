package comment

import "github.com/gin-gonic/gin"

type CommentHandler interface {
	PostCommentHdl(ctx *gin.Context)
}
