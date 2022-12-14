package photo

import (
	"context"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/errors"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/photo"
)

type PhotoUsecaseImpl struct {
	photoRepo photo.PhotoRepo
}

func NewPhotoUsecase(photoRepo photo.PhotoRepo) photo.PhotoUsecase {
	return &PhotoUsecaseImpl{photoRepo: photoRepo}
}

func (p *PhotoUsecaseImpl) PostPhotoSvc(ctx context.Context, userId uint, input photo.Photo) (photo.Photo, error) {
	if input.Title == "" {
		return photo.Photo{}, errors.ErrTitleEmpty
	}
	if input.Url == "" {
		return photo.Photo{}, errors.ErrUrlEmpty
	}
	err := p.photoRepo.PostPhoto(ctx, userId, &input)
	if err != nil {
		return photo.Photo{}, err
	}
	return input, nil
}

func (p *PhotoUsecaseImpl) GetPhotoByUseridSvc(ctx context.Context, userId uint) ([]photo.Photo, error) {
	result, err := p.photoRepo.GetPhotoByUserid(ctx, userId)
	if err != nil {
		return []photo.Photo{}, err
	}
	return result, nil
}

func (p *PhotoUsecaseImpl) UpdatePhotoSvc(ctx context.Context, userId, Id uint, input photo.Photo) (photo.Photo, error) {
	var result photo.Photo
	var err error
	result, err = p.photoRepo.UpdatePhoto(ctx, userId, Id, input)
	if err == errors.ErrUserNotFound {
		return photo.Photo{}, errors.ErrUserNotFound
	}
	if err != nil {
		return photo.Photo{}, err
	}
	return result, nil
}

func (p *PhotoUsecaseImpl) DeletePhotoSvc(ctx context.Context, userId, Id uint) error {
	err := p.photoRepo.DeletePhoto(ctx, userId, Id)

	if err != nil {
		return err
	}
	return nil
}
