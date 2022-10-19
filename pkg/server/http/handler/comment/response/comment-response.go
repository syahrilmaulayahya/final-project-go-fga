package response

import (
	"time"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/comment"
)

type PostCommentResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"date"`
}

func PostCommentResponseFromDomain(domain comment.Comment) PostCommentResponse {
	return PostCommentResponse{
		ID:        domain.ID,
		Message:   domain.Message,
		PhotoID:   domain.PhotoID,
		UserID:    domain.UserID,
		CreatedAt: domain.CreatedAt,
	}
}
