package user

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/claim"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/errors"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/message"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/usecase/crypto"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/user"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/handler/user/request"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/handler/user/response"
)

type UserHdlImpl struct {
	userUsecase user.UserUsecase
	middleware  crypto.Middleware
}

func NewUserHdl(userUseCase user.UserUsecase, middleware crypto.Middleware) user.UserHandler {
	return &UserHdlImpl{userUsecase: userUseCase, middleware: middleware}
}

func (u *UserHdlImpl) RegisterHdl(ctx *gin.Context) {
	var input request.UserRegisterReq
	if err := ctx.ShouldBind(&input); err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrBindPayload.Error(), errors.ErrBadRequest.Error())
		return
	}
	convertedInput := input.ToDomain()
	result, err := u.userUsecase.RegisterSvc(ctx, convertedInput)
	if err != nil {
		switch err {
		case errors.ErrWrongEmailFormat:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrWrongEmailFormatMsg.Error(), errors.ErrWrongEmailFormat.Error())
			return
		case errors.ErrEmailUsed:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrEmailUsedMsg.Error(), errors.ErrEmailUsed.Error())
			return
		case errors.ErrEmptyUsername:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrEmptyUsernameMsg.Error(), errors.ErrEmptyUsername.Error())
			return
		case errors.ErrUserNameUsed:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrUserNameUsedMsg.Error(), errors.ErrUserNameUsed.Error())
			return
		case errors.ErrWrongPasswordFormat:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrWrongPasswordFormatMsg.Error(), errors.ErrWrongPasswordFormat.Error())
			return
		case errors.ErrDobEmpty:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrDobEmptyMsg.Error(), errors.ErrDobEmpty.Error())
		case errors.ErrAgeRestriction:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrAgeRestrictionMsg.Error(), errors.ErrAgeRestriction.Error())
			return

		default:
			message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
			return
		}
	}
	response := response.UserRegisterFromDomain(result)

	message.SuccessResponseSwitcher(ctx, http.StatusAccepted, "user registered", response)
}

func (u *UserHdlImpl) GetUserByIdHdl(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("user_id"))

	result, err := u.userUsecase.GetUserByIdSvc(ctx, uint(id))
	if err != nil {
		switch err {
		case errors.ErrInternalServerError:
			message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
			return

		case errors.ErrUserNotFound:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrUserNotFoundMsg.Error(), errors.ErrUserNotFound.Error())
			return
		}
	}
	message.SuccessResponseSwitcher(ctx, http.StatusOK, "user found", response.GetUserByIdResponseFromDomain(result))
}

func (u *UserHdlImpl) UpdateUserHdl(ctx *gin.Context) {

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

	getClaim := u.middleware.VerifyJWT(ctx, bearerArray[1])
	result, err := u.userUsecase.GetUserByIdSvc(ctx, getClaim.Subject)
	if err != nil {
		switch err {
		case errors.ErrUserNotFound:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrUserNotFoundMsg.Error(), errors.ErrUserNotFound.Error())
			return

		default:
			message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
			return
		}
	}

	var input request.UserUpdate
	if err := ctx.ShouldBind(&input); err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrBindPayload.Error(), errors.ErrBadRequest.Error())
		return
	}

	result, err = u.userUsecase.UpdateUserSvc(ctx, input.Email, input.Username, int(result.ID))
	if err != nil {
		switch err {
		case errors.ErrWrongEmailFormat:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrWrongEmailFormatMsg.Error(), errors.ErrWrongEmailFormat.Error())
			return
		case errors.ErrEmailUsed:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrEmailUsedMsg.Error(), errors.ErrEmailUsed.Error())
			return
		case errors.ErrEmptyUsername:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrEmptyUsernameMsg.Error(), errors.ErrEmptyUsername.Error())
			return
		case errors.ErrUserNameUsed:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrUserNameUsedMsg.Error(), errors.ErrUserNameUsed.Error())
			return
		default:
			message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
			return
		}
	}

	claimIdToken := claim.ClaimIdToken{
		JWTID:    uuid.New(),
		Username: result.Username,
		Email:    result.Email,
		DOB:      result.Dob,
	}
	idToken, _ := u.middleware.CreateJWT(ctx, claimIdToken)
	response := map[string]string{
		"id_token": idToken,
	}
	message.SuccessResponseSwitcher(ctx, http.StatusAccepted, "updated", response)
}

func (u *UserHdlImpl) DeleteUserHdl(ctx *gin.Context) {
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

	getClaim := u.middleware.VerifyJWT(ctx, bearerArray[1])
	err := u.userUsecase.DeleteUserSvc(ctx, getClaim.Subject)
	if err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, "error while deleting data", errors.ErrInternalServerError.Error())
		return
	}
	message.SuccessResponseSwitcher(ctx, http.StatusOK, "your account has been succesfully deleted", nil)
}

func (u *UserHdlImpl) LoginHdl(ctx *gin.Context) {
	var input request.UserLogin
	if err := ctx.ShouldBind(&input); err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrBindPayload.Error(), errors.ErrBadRequest.Error())
		return
	}
	result, err := u.userUsecase.LoginSvc(ctx, input.Email, input.Password)
	if err != nil {
		switch err {
		case errors.ErrWrongEmailFormat:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrWrongEmailFormatMsg.Error(), errors.ErrWrongEmailFormat.Error())
			return
		case errors.ErrEmailNotFound:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrEmailNotFoundMsg.Error(), errors.ErrEmailNotFound.Error())
			return
		case errors.ErrWrongPassword:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrWrongPasswordMsg.Error(), errors.ErrWrongPassword.Error())
			return
		default:
			message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
			return
		}
	}
	timeNow := time.Now()
	claimAccess := claim.ClaimJwt{
		JWTID:          uuid.New(),
		Subject:        result.ID,
		Issuer:         "mygram.com",
		Audience:       "user.mygram.com.com",
		Scope:          "user",
		Type:           "ACCESS_TOKEN",
		IssuedAt:       timeNow.Unix(),
		NotValidBefore: timeNow.Unix(),
		ExpiredAt:      timeNow.Add(24 * time.Hour).Unix(),
	}
	accessToken, _ := u.middleware.CreateJWT(ctx, claimAccess)

	claimIdToken := claim.ClaimIdToken{
		JWTID:    uuid.New(),
		Username: result.Username,
		Email:    result.Email,
		DOB:      result.Dob,
	}
	idToken, _ := u.middleware.CreateJWT(ctx, claimIdToken)

	claimRefresh := claim.ClaimJwt{
		JWTID:          uuid.New(),
		Subject:        result.ID,
		Issuer:         "mygram.com",
		Audience:       "user.mygram.com.com",
		Scope:          "user",
		Type:           "REFRESH_TOKEN",
		IssuedAt:       timeNow.Unix(),
		NotValidBefore: timeNow.Unix(),
		ExpiredAt:      timeNow.Add(1000 * time.Hour).Unix(),
	}
	refreshToken, _ := u.middleware.CreateJWT(ctx, claimRefresh)

	response := response.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IdToken:      idToken,
	}

	message.SuccessResponseSwitcher(ctx, http.StatusAccepted, "logged in", response)
}

func (u *UserHdlImpl) RefreshHdl(ctx *gin.Context) {
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

	getClaim := u.middleware.VerifyJWT(ctx, bearerArray[1])
	result, err := u.userUsecase.GetUserByIdSvc(ctx, getClaim.Subject)
	if err != nil {
		switch err {
		case errors.ErrUserNotFound:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrUserNotFoundMsg.Error(), errors.ErrUserNotFound.Error())
			return
		default:
			message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
			return
		}
	}
	timeNow := time.Now()
	claimAccess := claim.ClaimJwt{
		JWTID:          uuid.New(),
		Subject:        result.ID,
		Issuer:         "mygram.com",
		Audience:       "user.mygram.com.com",
		Scope:          "user",
		Type:           "ACCESS_TOKEN",
		IssuedAt:       timeNow.Unix(),
		NotValidBefore: timeNow.Unix(),
		ExpiredAt:      timeNow.Add(24 * time.Hour).Unix(),
	}
	accessToken, _ := u.middleware.CreateJWT(ctx, claimAccess)

	claimIdToken := claim.ClaimIdToken{
		JWTID:    uuid.New(),
		Username: result.Username,
		Email:    result.Email,
		DOB:      result.Dob,
	}
	idToken, _ := u.middleware.CreateJWT(ctx, claimIdToken)

	claimRefresh := claim.ClaimJwt{
		JWTID:          uuid.New(),
		Subject:        result.ID,
		Issuer:         "mygram.com",
		Audience:       "user.mygram.com.com",
		Scope:          "user",
		Type:           "REFRESH_TOKEN",
		IssuedAt:       timeNow.Unix(),
		NotValidBefore: timeNow.Unix(),
		ExpiredAt:      timeNow.Add(1000 * time.Hour).Unix(),
	}
	refreshToken, _ := u.middleware.CreateJWT(ctx, claimRefresh)

	response := response.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IdToken:      idToken,
	}

	message.SuccessResponseSwitcher(ctx, http.StatusAccepted, "refresh token success", response)
}
