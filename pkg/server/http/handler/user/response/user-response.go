package response

import (
	"time"

	"github.com/syahrilmaulayahya/final-project-go-fga/helpers"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/user"
)

type UserRegisterResponse struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
}
type GetUserByIdResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func GetUserByIdResponseFromDomain(domain user.User) GetUserByIdResponse {
	return GetUserByIdResponse{
		ID:       domain.ID,
		Username: domain.Username,
	}
}

func UserRegisterFromDomain(domain user.User) UserRegisterResponse {
	return UserRegisterResponse{
		Age:      helpers.Age(domain.Dob, time.Now()),
		Email:    domain.Email,
		ID:       domain.ID,
		Username: domain.Username,
	}
}
