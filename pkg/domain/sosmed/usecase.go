package sosmed

import "context"

type SosmedUsecase interface {
	AddSosmedSvc(ctx context.Context, userId uint, input Sosmed) (Sosmed, error)
	GetSosmedByUserIdSvc(ctx context.Context, userId uint) ([]Sosmed, error)
	UpdateSosmedSvc(ctx context.Context, userId, Id uint, input Sosmed) (Sosmed, error)
	DeleteSosmedSvc(ctx context.Context, userId, Id uint) error
}
