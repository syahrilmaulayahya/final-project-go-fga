package sosmed

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/errors"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/message"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/sosmed"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/handler/sosmed/request"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/handler/sosmed/response"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/usecase/crypto"
)

type SosmedHdlImpl struct {
	sosmedUsecase sosmed.SosmedUsecase
	middleware    crypto.Middleware
}

func NewSomedHdl(sosmedUsecase sosmed.SosmedUsecase, middleware crypto.Middleware) sosmed.SosmedHandler {
	return &SosmedHdlImpl{sosmedUsecase: sosmedUsecase, middleware: middleware}
}

func (s *SosmedHdlImpl) AddSosmedHdl(ctx *gin.Context) {
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

	getClaim := s.middleware.VerifyJWT(ctx, bearerArray[1])
	var input request.AddSosmedRequest
	if err := ctx.ShouldBind(&input); err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrBindPayload.Error(), errors.ErrBadRequest.Error())
		return
	}
	result, err := s.sosmedUsecase.AddSosmedSvc(ctx, getClaim.Subject, input.ToDomain())
	if err != nil {
		switch err {
		case errors.ErrNameEmpty:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrNameEmptyMsg.Error(), errors.ErrNameEmpty.Error())
			return
		case errors.ErrUrlEmpty:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrUrlEmptyMsg.Error(), errors.ErrUrlEmpty.Error())
			return
		default:
			message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
			return
		}
	}
	message.SuccessResponseSwitcher(ctx, http.StatusAccepted, "add social media success", response.AddSosmedResponseFromDomain(result))
}

func (s *SosmedHdlImpl) GetSosmedByUserIdHdl(ctx *gin.Context) {
	var result []sosmed.Sosmed
	id, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrInvalidId.Error(), errors.ErrInvalidIdMsg.Error())
	}
	result, err = s.sosmedUsecase.GetSosmedByUserIdSvc(ctx, uint(id))
	if err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
		return
	}
	message.SuccessResponseSwitcher(ctx, http.StatusOK, "success getting data", response.ListGetSosmedResponseFromDomain(result))

}

func (s *SosmedHdlImpl) UpdateSosmedHdl(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("socialMediaId"))
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

	getClaim := s.middleware.VerifyJWT(ctx, bearerArray[1])
	var input request.AddSosmedRequest
	if err := ctx.ShouldBind(&input); err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrBindPayload.Error(), errors.ErrBadRequest.Error())
		return
	}
	result, err := s.sosmedUsecase.UpdateSosmedSvc(ctx, getClaim.Subject, uint(id), input.ToDomain())
	if err == errors.ErrSosmedNotFound {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrSosmedNotFoundMsg.Error(), errors.ErrSosmedNotFound.Error())
		return
	}
	if err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
		return
	}
	message.SuccessResponseSwitcher(ctx, http.StatusAccepted, "success update sosmed", response.UpdateSosmedResponseFromDomain(result))
}
