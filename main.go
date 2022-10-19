package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	engine "github.com/syahrilmaulayahya/final-project-go-fga/config/gin"
	"github.com/syahrilmaulayahya/final-project-go-fga/config/postgres"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/comment"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/message"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/photo"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/sosmed"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/user"
	userrepo "github.com/syahrilmaulayahya/final-project-go-fga/pkg/repository/user"
	userhandler "github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/handler/user"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/middlware"
	v1 "github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/router/v1"
	"gorm.io/gorm"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/usecase/crypto"
	userusecase "github.com/syahrilmaulayahya/final-project-go-fga/pkg/usecase/user"

	sosmedrepo "github.com/syahrilmaulayahya/final-project-go-fga/pkg/repository/sosmed"
	sosmedhandler "github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/handler/sosmed"
	sosmedusecase "github.com/syahrilmaulayahya/final-project-go-fga/pkg/usecase/sosmed"

	photorepo "github.com/syahrilmaulayahya/final-project-go-fga/pkg/repository/photo"
	photohandler "github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/handler/photo"
	photousecase "github.com/syahrilmaulayahya/final-project-go-fga/pkg/usecase/photo"

	commentrepo "github.com/syahrilmaulayahya/final-project-go-fga/pkg/repository/comment"
	commenthandler "github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/handler/comment"
	commentusecase "github.com/syahrilmaulayahya/final-project-go-fga/pkg/usecase/comment"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &sosmed.Sosmed{}, &photo.Photo{}, &comment.Comment{})
}

func init() {
	godotenv.Load(".env")

}
func main() {
	postgresHost := os.Getenv("MY_GRAM_POSTGRES_HOST")
	postgresPort := os.Getenv("MY_GRAM_POSTGRES_PORT")
	postgresDatabase := os.Getenv("MY_GRAM_POSTGRES_DATABASE")
	postgresUsername := os.Getenv("MY_GRAM_POSTGRES_USERNAME")
	postgresPassword := os.Getenv("MY_GRAM_POSTGRES_PASSWORD")
	sharedKey := os.Getenv("MY_GRAM_JWT_SHARED_KEY")

	postgresCln := postgres.NewPostgresConnection(postgres.Config{
		Host:         postgresHost,
		Port:         postgresPort,
		User:         postgresUsername,
		Password:     postgresPassword,
		DatabaseName: postgresDatabase,
	})
	dbMigrate(postgresCln.GetClient())
	ginEngine := engine.NewGinHttp(engine.Config{
		Port: ":8080",
	})
	middleware := crypto.NewMiddleware(crypto.Config{SharedKey: sharedKey})
	auth := middlware.NewAuth(middleware)
	ginEngine.GetGin().Use(
		gin.Recovery(),
		gin.Logger())

	starTime := time.Now()
	ginEngine.GetGin().GET("/", func(ctx *gin.Context) {
		response := message.SuccessResponse{}
		response.Code = 00
		response.Message = "server up and running"
		response.Type = "SUCCESS"
		response.StarTime = &starTime
		ctx.JSONP(http.StatusOK, response)
	})

	userRepo := userrepo.NewUserRepo(postgresCln)
	userUsecase := userusecase.NewUserUsecase(userRepo)
	userHandler := userhandler.NewUserHdl(userUsecase, middleware)

	sosmedRepo := sosmedrepo.NewSosmedRepo(postgresCln)
	sosmedUsecase := sosmedusecase.NewSosmedUsecase(sosmedRepo)
	sosmedHandler := sosmedhandler.NewSomedHdl(sosmedUsecase, middleware)

	photoRepo := photorepo.NewPhotoRepo(postgresCln)
	photoUsecase := photousecase.NewPhotoUsecase(photoRepo)
	photoHandler := photohandler.NewPhotoHdl(photoUsecase, middleware)

	commentRepo := commentrepo.NewCommentRepo(postgresCln)
	commentUsecase := commentusecase.NewCommentUsecase(commentRepo)
	commentHandler := commenthandler.NewCommentHandler(commentUsecase, middleware)

	v1.NewAuthRouter(ginEngine, userHandler, auth).Routers()
	v1.NewUserRouter(ginEngine, userHandler, auth).Routers()
	v1.NewSosmedRouter(ginEngine, sosmedHandler, auth).Routers()
	v1.NewPhotoRouter(ginEngine, photoHandler, auth).Routers()
	v1.NewCommentRouter(ginEngine, commentHandler, auth).Routers()

	ginEngine.Serve()
}
