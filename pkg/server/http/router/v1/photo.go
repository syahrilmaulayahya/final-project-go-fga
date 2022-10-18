package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/syahrilmaulayahya/final-project-go-fga/config/gin"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/photo"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/middlware"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/router"
)

type PhotoRouterImpl struct {
	ginEngine    engine.HttpServer
	routerGroup  *gin.RouterGroup
	photoHandler photo.PhotoHdl
	auth         middlware.Auth
}

func NewPhotoRouter(ginEngine engine.HttpServer, photoHandler photo.PhotoHdl, auth middlware.Auth) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/mygram/v1/photos")
	return &PhotoRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, photoHandler: photoHandler, auth: auth}
}

func (p *PhotoRouterImpl) post() {
	p.routerGroup.POST("/", p.auth.CheckJwt, p.photoHandler.PostPhotoHdl)

}
func (p *PhotoRouterImpl) get() {
	p.routerGroup.GET("/", p.auth.CheckJwt, p.photoHandler.GetOwnPhotoHdl)
	p.routerGroup.GET("/:user_id", p.photoHandler.GetPhotoByUserIdHdl)
}

func (p *PhotoRouterImpl) Routers() {
	p.post()
	p.get()
}
