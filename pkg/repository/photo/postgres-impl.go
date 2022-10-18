package photo

import (
	"context"

	"github.com/syahrilmaulayahya/final-project-go-fga/config/postgres"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/photo"
)

type PhotoRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewPhotoRepo(pgCln postgres.PostgresClient) photo.PhotoRepo {
	return &PhotoRepoImpl{pgCln: pgCln}
}

func (p *PhotoRepoImpl) PostPhoto(ctx context.Context, userId uint, input *photo.Photo) error {
	db := p.pgCln.GetClient()
	input.UserID = userId
	result := db.Model(&photo.Photo{}).Create(&input)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
