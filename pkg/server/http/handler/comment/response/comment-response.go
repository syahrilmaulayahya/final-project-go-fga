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
	CreatedAt time.Time `json:"created_at"`
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

type GetCommentResponse struct {
	ID        uint                 `json:"id"`
	Message   string               `json:"message"`
	PhotoID   uint                 `json:"photo_id"`
	UserID    uint                 `json:"user_id"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	Photo     PhotoCommentResponse `json:"photo"`
}

func GetCommentResponseFromDomain(domain comment.Comment) GetCommentResponse {
	return GetCommentResponse{
		ID:        domain.ID,
		Message:   domain.Message,
		PhotoID:   domain.PhotoID,
		UserID:    domain.UserID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		Photo:     PhotoCommentResponseFromDomain(domain.Photo),
	}
}

func ListGetCommentResponseFromDomain(domain []comment.Comment) []GetCommentResponse {
	var result []GetCommentResponse
	for _, comment := range domain {
		result = append(result, GetCommentResponseFromDomain(comment))
	}
	return result
}

type PhotoCommentResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Caption string `json:"caption"`
	Url     string `json:"url"`
	UserID  uint   `json:"user_id"`
}

func PhotoCommentResponseFromDomain(domain comment.Photo) PhotoCommentResponse {
	return PhotoCommentResponse{
		ID:      domain.ID,
		Title:   domain.Title,
		Caption: domain.Caption,
		Url:     domain.Url,
		UserID:  domain.UserID,
	}
}
