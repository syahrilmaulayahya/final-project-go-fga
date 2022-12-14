package photo

import "github.com/gin-gonic/gin"

type PhotoHdl interface {
	PostPhotoHdl(ctx *gin.Context)
	GetOwnPhotoHdl(ctx *gin.Context)
	GetPhotoByUserIdHdl(ctx *gin.Context)
	UpdatePhotoHdl(ctx *gin.Context)
	DeletePhotoHdl(ctx *gin.Context)
}
