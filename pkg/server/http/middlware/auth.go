package middlware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/errors"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/message"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/usecase/crypto"
)

type Auth interface {
	CheckJwt(ctx *gin.Context)
	CheckJwtRefresh(ctx *gin.Context)
}
type AuthImpl struct {
	middleware crypto.Middleware
}

func NewAuth(middleware crypto.Middleware) Auth {
	return &AuthImpl{middleware: middleware}
}

func (a *AuthImpl) CheckJwtRefresh(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")

	bearerArray := strings.Split(bearer, " ")

	if len(bearerArray) != 2 {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())
		return
	}

	if bearerArray[0] != "Bearer" {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())
		return
	}

	// claim := a.VerifyJWT(ctx, bearerArray[1])
	claim := a.middleware.VerifyJWT(ctx, bearerArray[1])

	// validate claim

	if claim.Issuer != "mygram.com" {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())
		return
	}

	if claim.Audience != "user.mygram.com.com" {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())

		return
	}

	if claim.Scope != "user" {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())
		return
	}
	if claim.Type != "REFRESH_TOKEN" {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, "please input refresh_token", errors.ErrUnauthorizhedReq.Error())
		return
	}
	if !time.Unix(claim.NotValidBefore, 0).Before(time.Now()) {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())
		return
	}

	if time.Unix(claim.ExpiredAt, 0).Before(time.Now()) {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())
		return
	}

	ctx.Set("user", claim.Subject)
	ctx.Next()
}
func (a *AuthImpl) CheckJwt(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")

	bearerArray := strings.Split(bearer, " ")

	if len(bearerArray) != 2 {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())
		return
	}

	if bearerArray[0] != "Bearer" {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())
		return
	}

	// claim := a.VerifyJWT(ctx, bearerArray[1])
	claim := a.middleware.VerifyJWT(ctx, bearerArray[1])

	// validate claim

	if claim.Issuer != "mygram.com" {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())
		return
	}

	if claim.Audience != "user.mygram.com.com" {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())

		return
	}

	if claim.Scope != "user" {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())
		return
	}
	if claim.Type != "ACCESS_TOKEN" {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, "please input access_token", errors.ErrUnauthorizhedReq.Error())
		return
	}
	if !time.Unix(claim.NotValidBefore, 0).Before(time.Now()) {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())
		return
	}

	if time.Unix(claim.ExpiredAt, 0).Before(time.Now()) {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized, errors.ErrUnauthorizhedReqMsg.Error(), errors.ErrUnauthorizhedReq.Error())
		return
	}

	ctx.Set("user", claim.Subject)
	ctx.Next()
}
