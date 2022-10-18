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

type GetPhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"url"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetPhotoResponseFromDomain(domain photo.Photo) GetPhotoResponse {
	return GetPhotoResponse{
		ID:        domain.ID,
		Title:     domain.Title,
		Caption:   domain.Caption,
		Url:       domain.Url,
		UserId:    domain.UserID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ListGetPhotoResponseFromDomain(domain []photo.Photo) []GetPhotoResponse {
	result := []GetPhotoResponse{}
	for _, photo := range domain {
		result = append(result, GetPhotoResponseFromDomain(photo))
	}
	return result
}

type UpdatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"url"`
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UpdatePhotoResponseFromDomain(domain photo.Photo) UpdatePhotoResponse {
	return UpdatePhotoResponse{
		ID:        domain.ID,
		Title:     domain.Title,
		Caption:   domain.Caption,
		Url:       domain.Url,
		UserId:    domain.UserID,
		UpdatedAt: domain.UpdatedAt,
	}
}
