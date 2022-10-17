package user

import (
	"context"
	"errors"

	"github.com/syahrilmaulayahya/final-project-go-fga/config/postgres"
	customError "github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/errors"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/user"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewUserRepo(pgCln postgres.PostgresClient) user.UserRepo {
	return &UserRepoImpl{pgCln: pgCln}
}

func (u *UserRepoImpl) Register(ctx context.Context, input *user.User) error {
	db := u.pgCln.GetClient()
	result := db.Model(&user.User{}).Create(&input)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepoImpl) GetUserById(ctx context.Context, id uint) (user.User, error) {
	var result user.User
	db := u.pgCln.GetClient()
	resultDb := db.Model(&user.User{}).First(&result, "id = ?", id)
	if errors.Is(resultDb.Error, gorm.ErrRecordNotFound) {
		return user.User{}, customError.ErrUserNotFound
	}
	if resultDb.Error != nil {
		return user.User{}, resultDb.Error
	}
	return result, nil
}

func (u *UserRepoImpl) UpdateUser(ctx context.Context, email, username string, id int) (user.User, error) {
	var result user.User

	db := u.pgCln.GetClient()
	resultDb := db.Model(&user.User{}).First(&result, "id = ? ", id)

	result.Email = email
	result.Username = username
	resultDb.Save(&result)
	if resultDb.Error != nil {
		return user.User{}, resultDb.Error
	}
	return result, nil
}

func (u *UserRepoImpl) DeleteUser(ctx context.Context, id uint) error {

	db := u.pgCln.GetClient()
	resultDb := db.Delete(&user.User{}, id)
	if resultDb.Error != nil {
		return resultDb.Error
	}
	return nil
}

func (u *UserRepoImpl) GetUserByUsername(ctx context.Context, username string) (user.User, error) {
	var result user.User
	db := u.pgCln.GetClient()
	resultDb := db.Model(&user.User{}).First(&result, "username = ?", username)
	if resultDb.Error != nil {
		return user.User{}, resultDb.Error
	}

	return result, nil
}

func (u *UserRepoImpl) GetUserByEmail(ctx context.Context, email string) (user.User, error) {
	var result user.User
	db := u.pgCln.GetClient()
	resultDb := db.Model(&user.User{}).First(&result, "email = ?", email)
	if resultDb.Error != nil {
		return user.User{}, resultDb.Error
	}
	return result, nil
}

func (u *UserRepoImpl) Login(ctx context.Context, email, password string) (user.User, error) {
	var result user.User

	db := u.pgCln.GetClient()
	resultDb := db.Model(&user.User{}).First(&result, "email = ? ", email)
	if resultDb.Error != nil {
		return user.User{}, resultDb.Error
	}

	return result, nil
}
