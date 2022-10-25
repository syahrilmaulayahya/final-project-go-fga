package crypto

import (
	"context"

	"github.com/kataras/jwt"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/claim"
)

type Middleware interface {
	CreateJWT(ctx context.Context, claim any) (string, error)
	VerifyJWT(ctx context.Context, token string) (claims claim.ClaimJwt)
	VerifyClaimIdToken(ctx context.Context, token string) (claims claim.ClaimIdToken)
}

type Config struct {
	SharedKey string
}

type MiddlewareImpl struct {
	config Config
}

func NewMiddleware(config Config) Middleware {
	return &MiddlewareImpl{config: config}
}

func (m *MiddlewareImpl) CreateJWT(ctx context.Context, claim any) (string, error) {

	token, err := jwt.Sign(jwt.HS256, []byte(m.config.SharedKey), claim)
	if err != nil {
		return "", err
	}
	return string(token), nil

}
func (m *MiddlewareImpl) VerifyJWT(ctx context.Context, token string) (claims claim.ClaimJwt) {

	verifiedToken, err := jwt.Verify(jwt.HS256, []byte(m.config.SharedKey), []byte(token))
	if err != nil {
		panic(err)
	}

	err = verifiedToken.Claims(&claims)
	if err != nil {
		panic(err)
	}
	return claims

}
func (m *MiddlewareImpl) VerifyClaimIdToken(ctx context.Context, token string) (claims claim.ClaimIdToken) {

	verifiedToken, err := jwt.Verify(jwt.HS256, m.config.SharedKey, []byte(token))
	if err != nil {
		panic(err)
	}

	err = verifiedToken.Claims(&claims)
	if err != nil {
		panic(err)
	}
	return claims
}

// var (
// 	sharedKey = os.Getenv("MY_GRAM_JWT_SHARED_KEY")
// )

// func CreateJWT(ctx context.Context, claim any) (string, error) {
// 	godotenv.Load(".env")
// 	token, err := jwt.Sign(jwt.HS256, sharedKey, claim)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(token), nil
// }

// func VerifyJWT(ctx context.Context, token string) (claims claim.ClaimJwt) {
// 	godotenv.Load(".env")

// 	verifiedToken, err := jwt.Verify(jwt.HS256, sharedKey, []byte(token))
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = verifiedToken.Claims(&claims)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return claims
// }

// func VerifyClaimIdToken(ctx context.Context, token string) (claims claim.ClaimIdToken) {
// 	godotenv.Load(".env")

// 	verifiedToken, err := jwt.Verify(jwt.HS256, sharedKey, []byte(token))
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = verifiedToken.Claims(&claims)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return claims
// }
