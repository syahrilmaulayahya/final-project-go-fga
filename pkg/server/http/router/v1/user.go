package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/syahrilmaulayahya/final-project-go-fga/config/gin"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/user"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/router"
)

type UserRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	userHandler user.UserHandler
}

func NewUserRouter(ginEngine engine.HttpServer, userHandler user.UserHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/mygram/v1/users")
	return &UserRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, userHandler: userHandler}
}

func (u *UserRouterImpl) post() {
	u.routerGroup.POST("/register", u.userHandler.RegisterHdl)
	u.routerGroup.PUT("/update", u.userHandler.UpdateUserHdl)

}
func (u *UserRouterImpl) get() {
	u.routerGroup.GET("/get", u.userHandler.GetUserByIdHdl)

}
func (u *UserRouterImpl) Routers() {
	u.post()
	u.get()
}
