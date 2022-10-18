package photo

import "github.com/gin-gonic/gin"

type PhotoHdl interface {
	PostPhotoHdl(ctx *gin.Context)
}
