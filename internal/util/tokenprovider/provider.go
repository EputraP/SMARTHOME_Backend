package tokenprovider

import (
	"log"
	"time"

	"github.com/EputraP/SMARTHOME_Backend/internal/model"
	"github.com/golang-jwt/jwt/v4"
)

type JWTTokenProvider interface {
	GenerateRefreshToken(user model.Account) (string, error)
	GenerateAccessToken(user model.Account)(string, error)
	ValidateToken(token string)(*JwtClaims, error)
}

type jwtTokenProvider struct {
	issuer               string
	secret               string
	refreshTokenDuration int
	accessTokenDuration  int
}

func NewJWT(issuer string, secret string, refreshTokenDuration int, accessTokenDuration int) JWTTokenProvider {
	return &jwtTokenProvider{
		issuer:               issuer,
		secret:               secret,
		refreshTokenDuration: refreshTokenDuration,
		accessTokenDuration:  accessTokenDuration,
	}
}

func (p *jwtTokenProvider) GenerateAccessToken(user model.User) (string, error) {
	return p.generateToken(user, time.Duration(p.accessTokenDuration)*time.Minute)
}

func (p *jwtTokenProvider) GenerateRefreshToken(user model.User) (string, error) {
	return p.generateToken(user, time.Duration(p.refreshTokenDuration)*time.Minute)
}

func (p *jwtTokenProvider) generateToken(user model.User, expiresIn time.Duration) (string, error) {
	claims := JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    p.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserClaims: UserClaims{
			UserID:   user.UserId.String(),
			Username: user.Username,
			Role:     user.Role,
			Type:     constant.TypeUserClaim,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(p.secret))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenStr, nil
}

func (p *jwtTokenProvider) ValidateToken(token string) (*JwtClaims, error) {
	claims := JwtClaims{}

	jwtToken, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(p.secret), nil
	})

	if jwtToken == nil || !jwtToken.Valid {
		return nil, errs.InvalidToken
	}

	if err != nil {
		return nil, err
	}

	if claims.Issuer != p.issuer {
		return nil, errs.InvalidIssuer
	}

	return &claims, nil
}



