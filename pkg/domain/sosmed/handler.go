package sosmed

import (
	"github.com/gin-gonic/gin"
)

type SosmedHandler interface {
	AddSosmedHdl(ctx *gin.Context)
}
