package sosmed

import "context"

type SosmedRepo interface {
	AddSosmed(ctx context.Context, userId uint, input *Sosmed) error
	GetSosmedByUserId(ctx context.Context, userId uint) ([]Sosmed, error)
	UpdateSosmed(ctx context.Context, userId, Id uint, input Sosmed) (Sosmed, error)
	DeleteSosmed(ctx context.Context, userId, Id uint) error
}
