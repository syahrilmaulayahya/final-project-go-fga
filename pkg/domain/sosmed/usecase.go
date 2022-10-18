package sosmed

import "context"

type SosmedUsecase interface {
	AddSosmedSvc(ctx context.Context, userId uint, input Sosmed) (Sosmed, error)
	GetSosmedByUserIdSvc(ctx context.Context, userId uint) ([]Sosmed, error)
}
