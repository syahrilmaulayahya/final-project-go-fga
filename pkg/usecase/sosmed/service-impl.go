package sosmed

import (
	"context"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/errors"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/sosmed"
)

type SosmedUsecaseImpl struct {
	sosmedRepo sosmed.SosmedRepo
}

func NewSosmedUsecase(sosmedRepo sosmed.SosmedRepo) sosmed.SosmedUsecase {
	return &SosmedUsecaseImpl{sosmedRepo: sosmedRepo}
}

func (s *SosmedUsecaseImpl) AddSosmedSvc(ctx context.Context, userId uint, input sosmed.Sosmed) (sosmed.Sosmed, error) {
	if input.Name == "" {
		return sosmed.Sosmed{}, errors.ErrNameEmpty
	}
	if input.Url == "" {
		return sosmed.Sosmed{}, errors.ErrUrlEmpty
	}
	err := s.sosmedRepo.AddSosmed(ctx, userId, &input)
	if err != nil {
		return sosmed.Sosmed{}, err
	}
	return input, nil
}

func (s *SosmedUsecaseImpl) GetSosmedByUserIdSvc(ctx context.Context, userId uint) ([]sosmed.Sosmed, error) {
	var result []sosmed.Sosmed
	var err error
	result, err = s.sosmedRepo.GetSosmedByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
