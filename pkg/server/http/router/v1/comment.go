package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/syahrilmaulayahya/final-project-go-fga/config/gin"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/comment"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/middlware"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/router"
)

type CommentRouterImpl struct {
	ginEngine      engine.HttpServer
	routerGroup    *gin.RouterGroup
	commentHandler comment.CommentHandler
	auth           middlware.Auth
}

func NewCommentRouter(ginEngine engine.HttpServer, commentHandler comment.CommentHandler, auth middlware.Auth) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/mygram/v1/comments")
	return &CommentRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, commentHandler: commentHandler, auth: auth}
}

func (c *CommentRouterImpl) post() {
	c.routerGroup.POST("/", c.auth.CheckJwt, c.commentHandler.PostCommentHdl)

}
func (c *CommentRouterImpl) get() {
	c.routerGroup.GET("/:userId", c.auth.CheckJwt, c.commentHandler.GetCommentByUserIdHdl)

}
func (c *CommentRouterImpl) Routers() {
	c.post()
	c.get()

}
