package request

import (
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/photo"
)

type PostPhotoRequest struct {
	Title   string `json:"title"`
	Caption string `json:"caption"`
	Url     string `json:"url"`
}

func (p *PostPhotoRequest) ToDomain() photo.Photo {
	return photo.Photo{
		Title:   p.Title,
		Caption: p.Caption,
		Url:     p.Url,
	}
}
