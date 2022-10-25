package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/syahrilmaulayahya/final-project-go-fga/config/gin"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/user"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/middlware"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/router"
)

type AuthRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	userHandler user.UserHandler
	auth        middlware.Auth
}

func NewAuthRouter(ginEngine engine.HttpServer, userHandler user.UserHandler, auth middlware.Auth) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/mygram/v1/auth")
	return &AuthRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, userHandler: userHandler, auth: auth}
}

func (u *AuthRouterImpl) post() {
	u.routerGroup.POST("/login", u.userHandler.LoginHdl)
	u.routerGroup.POST("/refresh", u.auth.CheckJwtRefresh, u.userHandler.RefreshHdl)
}
func (u *AuthRouterImpl) delete() {

	u.routerGroup.DELETE("/users", u.auth.CheckJwt, u.userHandler.DeleteUserHdl)
}

func (u *AuthRouterImpl) Routers() {
	u.post()
	u.delete()
}
