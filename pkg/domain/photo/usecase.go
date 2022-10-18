package photo

import "context"

type PhotoUsecase interface {
	PostPhotoSvc(ctx context.Context, userId uint, input Photo) (Photo, error)
}
