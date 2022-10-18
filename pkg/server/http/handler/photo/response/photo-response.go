package response

import (
	"time"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/photo"
)

type PostPhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"url"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func PostPhotoResponseFromDomain(domain photo.Photo) PostPhotoResponse {
	return PostPhotoResponse{
		ID:        domain.ID,
		Title:     domain.Title,
		Caption:   domain.Caption,
		Url:       domain.Url,
		UserId:    domain.UserID,
		CreatedAt: domain.CreatedAt,
	}
}
