package photo

import "context"

type PhotoUsecase interface {
	PostPhotoSvc(ctx context.Context, userId uint, input Photo) (Photo, error)
	GetPhotoByUseridSvc(ctx context.Context, userId uint) ([]Photo, error)
}
