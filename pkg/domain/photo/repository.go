package photo

import "context"

type PhotoRepo interface {
	PostPhoto(ctx context.Context, userId uint, input *Photo) error
	GetPhotoByUserid(ctx context.Context, userId uint) ([]Photo, error)
}
