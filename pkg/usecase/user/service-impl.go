package user

import (
	"context"
	"time"

	customError "github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/errors"

	"github.com/syahrilmaulayahya/final-project-go-fga/helpers"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func NewUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}

func (u *UserUsecaseImpl) RegisterSvc(ctx context.Context, input user.User) (user.User, error) {
	var err error
	var result user.User

	// email validation
	if !helpers.CheckEmail(input.Email) {
		return user.User{}, customError.ErrWrongEmailFormat
	}

	result, _ = u.userRepo.GetUserByEmail(ctx, input.Email)

	if result.ID > 0 {
		err = customError.ErrEmailUsed
		return user.User{}, err
	}

	// username validation
	if input.Username == "" {
		return user.User{}, customError.ErrEmptyUsername
	}
	result, _ = u.userRepo.GetUserByUsername(ctx, input.Username)

	if result.ID > 0 {
		err = customError.ErrUserNameUsed
		return user.User{}, err
	}

	// password validation
	if len(input.Password) < 6 {
		return user.User{}, customError.ErrWrongPasswordFormat
	}
	passwordByte, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(passwordByte)

	// age validation
	if input.Dob.IsZero() {
		return user.User{}, customError.ErrDobEmpty
	}

	if helpers.Age(input.Dob, time.Now()) < 8 {
		return user.User{}, customError.ErrAgeRestriction
	}

	err = u.userRepo.Register(ctx, &input)
	if err != nil {
		return user.User{}, err
	}
	return input, nil
}

func (u *UserUsecaseImpl) GetUserByIdSvc(ctx context.Context, id uint) (user.User, error) {

	result, err := u.userRepo.GetUserById(ctx, id)
	if err == customError.ErrUserNotFound {
		return user.User{}, customError.ErrUserNotFound
	}
	if err != nil {
		return user.User{}, customError.ErrInternalServerError
	}

	return result, nil
}

func (u *UserUsecaseImpl) UpdateUserSvc(ctx context.Context, email, username string, id int) (user.User, error) {

	var err error
	var result user.User

	// email validation
	if !helpers.CheckEmail(email) {
		return user.User{}, customError.ErrWrongEmailFormat
	}

	result, _ = u.userRepo.GetUserByEmail(ctx, email)

	if result.ID > 0 {
		err = customError.ErrEmailUsed
		return user.User{}, err
	}

	// username validation
	if username == "" {
		return user.User{}, customError.ErrEmptyUsername
	}
	result, _ = u.userRepo.GetUserByUsername(ctx, username)

	if result.ID > 0 {
		err = customError.ErrUserNameUsed
		return user.User{}, err
	}
	result, err = u.userRepo.UpdateUser(ctx, email, username, id)
	if err != nil {
		return user.User{}, nil
	}
	return result, nil
}

func (u *UserUsecaseImpl) DeleteUserSvc(ctx context.Context, id uint) error {
	err := u.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return err
}
func (u *UserUsecaseImpl) LoginSvc(ctx context.Context, email, password string) (user.User, error) {
	// email validation
	var result user.User
	var err error
	if !helpers.CheckEmail(email) {
		return user.User{}, customError.ErrWrongEmailFormat
	}

	result, _ = u.userRepo.GetUserByEmail(ctx, email)

	if result.ID <= 0 {
		err = customError.ErrEmailNotFound
		return user.User{}, err
	}
	// password validation
	result, err = u.userRepo.Login(ctx, email, password)

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		return user.User{}, customError.ErrWrongPassword
	}
	if err != nil {
		return user.User{}, err
	}
	return result, nil
}

func (u *UserUsecaseImpl) GetUserByUsernameSvc(ctx context.Context, username string) (user.User, error) {
	// email validation
	var err error
	var result user.User

	// username validation
	if username == "" {
		return user.User{}, customError.ErrEmptyUsername
	}
	result, err = u.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return user.User{}, customError.ErrInternalServerError
	}
	if result.ID <= 0 {
		err = customError.ErrUsernameNotFound
		return user.User{}, err
	}

	return result, nil
}
