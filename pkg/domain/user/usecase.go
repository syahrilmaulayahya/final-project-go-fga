package user

import "context"

type UserUsecase interface {
	RegisterSvc(ctx context.Context, input User) (User, error)
	GetUserByIdSvc(ctx context.Context, id uint) (User, error)
	UpdateUserSvc(ctx context.Context, email, username string, id int) (User, error)
	DeleteUserSvc(ctx context.Context, id uint) error
	LoginSvc(ctx context.Context, email, password string) (User, error)
	GetUserByUsernameSvc(ctx context.Context, username string) (User, error)
}
