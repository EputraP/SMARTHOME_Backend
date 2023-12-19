package tokenprovider

import "github.com/golang-jwt/jwt/v4"

type UserClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Type     string `json:"type"`
}

type JwtClaims struct {
	jwt.RegisteredClaims
	UserClaims
}
