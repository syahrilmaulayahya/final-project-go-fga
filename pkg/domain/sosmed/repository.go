package sosmed

import "context"

type SosmedRepo interface {
	AddSosmed(ctx context.Context, userId uint, input *Sosmed) error
	GetSosmedByUserId(ctx context.Context, userId uint) ([]Sosmed, error)
}
