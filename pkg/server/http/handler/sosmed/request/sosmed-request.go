package request

import "github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/sosmed"

type AddSosmedRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (a *AddSosmedRequest) ToDomain() sosmed.Sosmed {
	return sosmed.Sosmed{
		Name: a.Name,
		Url:  a.Url,
	}
}
