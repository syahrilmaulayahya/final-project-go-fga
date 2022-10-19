package comment

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/comment"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/errors"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/message"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/handler/comment/request"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/server/http/handler/comment/response"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/usecase/crypto"
)

type CommentHandlerImpl struct {
	commentUsecase comment.CommentUsecase
	middleware     crypto.Middleware
}

func NewCommentHandler(commentUsecase comment.CommentUsecase, middleware crypto.Middleware) comment.CommentHandler {
	return &CommentHandlerImpl{commentUsecase: commentUsecase, middleware: middleware}
}

func (c *CommentHandlerImpl) PostCommentHdl(ctx *gin.Context) {
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

	getClaim := c.middleware.VerifyJWT(ctx, bearerArray[1])
	var input request.PostCommentRequest
	if err := ctx.ShouldBind(&input); err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrBindPayload.Error(), errors.ErrBadRequest.Error())
		return
	}
	input.UserId = getClaim.Subject
	result, err := c.commentUsecase.PostCommentSvc(ctx, request.PostCommentToDomain(input))

	if err != nil {
		switch err {
		case errors.ErrMessageEmpty:
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, errors.ErrMessageEmptyMsg.Error(), errors.ErrMessageEmpty.Error())
			return
		default:
			message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError, errors.ErrInternalServerErrorMsg.Error(), errors.ErrInternalServerError.Error())
			return
		}
	}
	message.SuccessResponseSwitcher(ctx, http.StatusAccepted, "post comment success", (response.PostCommentResponseFromDomain(result)))

}
