package sosmed

import (
	"github.com/gin-gonic/gin"
)

type SosmedHandler interface {
	AddSosmedHdl(ctx *gin.Context)
	GetSosmedByUserIdHdl(ctx *gin.Context)
	UpdateSosmedHdl(ctx *gin.Context)
	DeleteSosmedHdl(ctx *gin.Context)
}
