package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/syahrilmaulayahya/final-project-go-fga/config/gin"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/sosmed"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/middlware"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/router"
)

type SosmedRouterImpl struct {
	ginEngine     engine.HttpServer
	routerGroup   *gin.RouterGroup
	sosmedHandler sosmed.SosmedHandler
	auth          middlware.Auth
}

func NewSosmedRouter(ginEngine engine.HttpServer, sosmedHandler sosmed.SosmedHandler, auth middlware.Auth) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/mygram/v1/sosmed")
	return &SosmedRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, sosmedHandler: sosmedHandler, auth: auth}
}

func (s *SosmedRouterImpl) post() {
	s.routerGroup.POST("/", s.auth.CheckJwt, s.sosmedHandler.AddSosmedHdl)

}
func (s *SosmedRouterImpl) get() {
	s.routerGroup.GET("/:userId", s.sosmedHandler.GetSosmedByUserIdHdl)

}

func (s *SosmedRouterImpl) Routers() {
	s.post()
	s.get()
}
