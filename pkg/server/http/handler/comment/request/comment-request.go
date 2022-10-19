package request

import "github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/comment"

type PostCommentRequest struct {
	Message string `json:"message"`
	PhotoID uint   `json:"photo_id"`
	UserId  uint   `json:"-"`
}

func PostCommentToDomain(p PostCommentRequest) comment.Comment {
	return comment.Comment{
		Message: p.Message,
		PhotoID: p.PhotoID,
		UserID:  p.UserId,
	}
}
