package photo

import (
	"context"
	"errors"

	customError "github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/errors"

	"github.com/syahrilmaulayahya/final-project-go-fga/config/postgres"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/photo"
	"gorm.io/gorm"
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

func (p *PhotoRepoImpl) GetPhotoByUserid(ctx context.Context, userId uint) ([]photo.Photo, error) {
	var result []photo.Photo
	db := p.pgCln.GetClient()
	err := db.Preload("Comments").Find(&result, "user_id = ?", userId)
	if err.Error != nil {
		return nil, err.Error
	}
	return result, nil
}

func (p *PhotoRepoImpl) UpdatePhoto(ctx context.Context, userId, id uint, input photo.Photo) (photo.Photo, error) {
	var result photo.Photo
	db := p.pgCln.GetClient()
	err := db.Model(&photo.Photo{}).First(&result, "user_id = ? and id = ?", userId, id)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return photo.Photo{}, customError.ErrPhotoNotFound
	}
	if input.Title != "" {
		result.Title = input.Title
	}
	result.Caption = input.Caption
	if input.Url != "" {
		result.Url = input.Url
	}

	err.Save(&result)
	if err.Error != nil {
		return photo.Photo{}, err.Error
	}

	return result, nil
}

func (p *PhotoRepoImpl) DeletePhoto(ctx context.Context, userId, Id uint) error {
	var result photo.Photo
	db := p.pgCln.GetClient()
	resultDb := db.Model(&photo.Photo{}).First(&result, "user_id = ? and id = ?", userId, Id)

	if errors.Is(resultDb.Error, gorm.ErrRecordNotFound) {
		return customError.ErrPhotoNotFound
	}
	if resultDb.Error != nil {
		return resultDb.Error
	}
	resultDb = db.Delete(&result)

	if resultDb.Error != nil {
		return resultDb.Error
	}

	return nil
}
