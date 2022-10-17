package user

import "github.com/gin-gonic/gin"

type UserHandler interface {
	RegisterHdl(ctx *gin.Context)
	GetUserByIdHdl(ctx *gin.Context)
	UpdateUserHdl(ctx *gin.Context)
	DeleteUserHdl(ctx *gin.Context)
	LoginHdl(ctx *gin.Context)
	RefreshHdl(ctx *gin.Context)
}
