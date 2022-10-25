package photo

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/errors"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/message"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/photo"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/handler/photo/request"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/handler/photo/response"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/usecase/crypto"
)

type PhotoHdlImpl struct {
	photoUseCase photo.PhotoUsecase
	middleware   crypto.Middleware
}

func NewPhotoHdl(photoUsecase photo.PhotoUsecase, middleware crypto.Middleware) photo.PhotoHdl {
	return &PhotoHdlImpl{photoUseCase: photoUsecase, middleware: middleware}
}

func (p *PhotoHdlImpl) PostPhotoHdl(ctx *gin.Context) {
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

	getClaim := p.middleware.VerifyJWT(ctx, bearerArray[1])
	var input request.PostPhotoRequest
	if err := ctx.ShouldBind(&input); err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrBindPayload.Error(), errors.ErrBadRequest.Error())
		return
	}
	result, err := p.photoUseCase.PostPhotoSvc(ctx, getClaim.Subject, input.ToDomain())
	if err != nil {
		switch err {
		case errors.ErrTitleEmpty:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrTitleEmptyMsg.Error(), errors.ErrTitleEmpty.Error())
			return
		case errors.ErrUrlEmpty:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrUrlEmptyMsg.Error(), errors.ErrUrlEmpty.Error())
			return
		default:
			message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
			return
		}
	}
	message.SuccessResponseSwitcher(ctx, http.StatusAccepted, "post photo success", (response.PostPhotoResponseFromDomain(result)))
}

func (p *PhotoHdlImpl) GetOwnPhotoHdl(ctx *gin.Context) {
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

	getClaim := p.middleware.VerifyJWT(ctx, bearerArray[1])

	result, err := p.photoUseCase.GetPhotoByUseridSvc(ctx, getClaim.Subject)
	if err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
		return
	}
	message.SuccessResponseSwitcher(ctx, http.StatusOK, "get photo success", (response.ListGetPhotoResponseFromDomain(result)))

}

func (p *PhotoHdlImpl) GetPhotoByUserIdHdl(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrInvalidId.Error(), errors.ErrInvalidIdMsg.Error())
	}
	result, err := p.photoUseCase.GetPhotoByUseridSvc(ctx, uint(id))
	if err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
		return
	}
	message.SuccessResponseSwitcher(ctx, http.StatusOK, "get photo success", (response.ListGetPhotoResponseFromDomain(result)))
}

func (p *PhotoHdlImpl) UpdatePhotoHdl(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrInvalidId.Error(), errors.ErrInvalidIdMsg.Error())
	}
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

	getClaim := p.middleware.VerifyJWT(ctx, bearerArray[1])
	var input request.PostPhotoRequest
	if err := ctx.ShouldBind(&input); err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrBindPayload.Error(), errors.ErrBadRequest.Error())
		return
	}
	result, err := p.photoUseCase.UpdatePhotoSvc(ctx, getClaim.Subject, uint(id), input.ToDomain())
	if err == errors.ErrPhotoNotFound {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrPhotoNotFoundMsg.Error(), errors.ErrPhotoNotFound.Error())
		return
	}
	if err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
		return
	}
	message.SuccessResponseSwitcher(ctx, http.StatusAccepted, "success update photo", response.UpdatePhotoResponseFromDomain(result))
}

func (p *PhotoHdlImpl) DeletePhotoHdl(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrInvalidId.Error(), errors.ErrInvalidIdMsg.Error())
	}
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

	getClaim := p.middleware.VerifyJWT(ctx, bearerArray[1])

	err = p.photoUseCase.DeletePhotoSvc(ctx, getClaim.Subject, uint(id))
	if err == errors.ErrPhotoNotFound {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrPhotoNotFoundMsg.Error(), errors.ErrPhotoNotFound.Error())
		return
	}
	if err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
		return
	}
	message.SuccessResponseSwitcher(ctx, http.StatusOK, "your photo has been successfully deleted", nil)

}
