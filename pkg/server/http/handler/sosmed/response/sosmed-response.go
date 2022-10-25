package response

import (
	"time"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/sosmed"
)

type AddSosmedResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func AddSosmedResponseFromDomain(domain sosmed.Sosmed) AddSosmedResponse {
	return AddSosmedResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		Url:       domain.Url,
		UserId:    domain.UserID,
		CreatedAt: domain.CreatedAt,
	}
}

type UpdateSosmedResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"udpated_at"`
}

func UpdateSosmedResponseFromDomain(domain sosmed.Sosmed) UpdateSosmedResponse {
	return UpdateSosmedResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		Url:       domain.Url,
		UserId:    domain.UserID,
		UpdatedAt: domain.UpdatedAt,
	}
}

type GetSosmedResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetSosmedResponseFromDomain(domain sosmed.Sosmed) GetSosmedResponse {
	return GetSosmedResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		Url:       domain.Url,
		UserId:    domain.UserID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ListGetSosmedResponseFromDomain(domain []sosmed.Sosmed) []GetSosmedResponse {
	result := []GetSosmedResponse{}
	for _, sosmed := range domain {
		result = append(result, GetSosmedResponseFromDomain(sosmed))
	}
	return result
}
