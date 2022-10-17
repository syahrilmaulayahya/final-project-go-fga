package user

import "context"

type UserRepo interface {
	Register(ctx context.Context, input *User) error
	GetUserById(ctx context.Context, id uint) (User, error)
	UpdateUser(ctx context.Context, email, username string, id int) (User, error)
	DeleteUser(ctx context.Context, id uint) error
	Login(ctx context.Context, email, password string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
}
