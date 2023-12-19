package service

import (
	"regexp"
	"strings"

	"github.com/EputraP/SMARTHOME_Backend/internal/dto"
	"github.com/EputraP/SMARTHOME_Backend/internal/errs"
	"github.com/EputraP/SMARTHOME_Backend/internal/repository"
	"github.com/EputraP/SMARTHOME_Backend/internal/util/hasher"
	"github.com/EputraP/SMARTHOME_Backend/internal/util/tokenprovider"
	"gorm.io/gorm"
)

type AuthService interface {
	// CreateUser(input dto.RegisterBody) (*dto.RegisterResponse, error)
	// Login(input dto.LoginBody) (*dto.LoginResponse, error)
}

type authService struct {
	authRepo    repository.AuthRepository
	hasher      hasher.Hasher
	jtwProvider tokenprovider.JWTTokenProvider
}

type AuthServiceConfig struct {
	AuthRepo    repository.AuthRepository
	Hasher      hasher.Hasher
	JwtProvider tokenprovider.JWTTokenProvider
}

func NewAuthService(config AuthServiceConfig) AuthService {
	return &authService{
		authRepo:    config.AuthRepo,
		hasher:      config.Hasher,
		jtwProvider: config.JwtProvider,
	}
}

func (as authService) CreateUser(input dto.RegisterBody) (*dto.RegisterResponse, error) {
	lowerUsername :=  strings.ToLower(input.UserName) 

	_, err := as.authRepo.SearchUser(lowerUsername)

	if err == nil {
		return nil, errs.UsernameAlreadyUsed
	}

	re := regexp.MustCompile(`(?i)` + input.UserName)
	isMatch := re.MatchString(input.Password)

	if isMatch {
		return nil, errs.PasswordContainUsername
	}

	resp := &dto.RegisterResponse{}

	err = repository.AsTransaction(func(tx *gorm.DB) error {
		repoWithTx := as.authRepo.WithTx(tx)

		hashedPassword, _ := as.hasher.Hash(input.Password)

		registerInput := dto.RegisterBody{
			UserName: lowerUsername,
			Password: hashedPassword,
		}

		newUser, err := repoWithTx.CreateUser(registerInput)

		if err != nil {
			return err
		}

		resp = &dto.RegisterResponse{
			UserID:   newUser.ID,
			Username: newUser.Username,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}