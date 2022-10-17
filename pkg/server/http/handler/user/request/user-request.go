package request

import (
	"time"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/user"
)

type UserRegisterReq struct {
	Dob      string `json:"dob"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserUpdate struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func parseDob(dob string) time.Time {
	parseDob, _ := time.Parse("2006-Jan-02", dob)
	return parseDob
}

func (u *UserRegisterReq) ToDomain() user.User {
	return user.User{
		Dob:      parseDob(u.Dob),
		Email:    u.Email,
		Password: u.Password,
		Username: u.Username,
	}
}
