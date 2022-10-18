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
