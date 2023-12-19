package handler

import (
	"errors"

	"github.com/EputraP/SMARTHOME_Backend/internal/dto"
	"github.com/EputraP/SMARTHOME_Backend/internal/errs"
	"github.com/EputraP/SMARTHOME_Backend/internal/service"
	"github.com/EputraP/SMARTHOME_Backend/internal/util/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService service.AuthService
}

type AuthHandlerConfig struct {
	AuthService service.AuthService
}

func NewAuthHandler(config AuthHandlerConfig) *AuthHandler {
	return &AuthHandler{
		authService: config.AuthService,
	}
}

func (h AuthHandler) CreateUser(c *gin.Context) {
	var registerBody dto.RegisterBody

	if err := c.ShouldBindJSON(&registerBody); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.authService.CreateUser(registerBody)

	if err != nil {
		if errors.Is(err, errs.UsernameAlreadyUsed) ||
			errors.Is(err, errs.PasswordContainUsername) {
			response.Error(c, 400, err.Error())
			return
		}

		response.UnknownError(c, err)
		return
	}

	response.JSON(c, 201, "Register Success", resp)
}

func (h AuthHandler) Login(c *gin.Context) {
	var loginBody dto.LoginBody

	if err := c.ShouldBindJSON(&loginBody); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.authService.Login(loginBody)

	if err != nil {
		if errors.Is(err, errs.PasswordDoesntMatch) ||
			errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, 401, errs.UsernamePasswordIncorrect.Error())
			return
		}

		response.UnknownError(c, err)
		return
	}

	response.JSON(c, 200, "Login success", resp)
}