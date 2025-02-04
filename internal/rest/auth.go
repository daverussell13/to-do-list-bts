package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	Login()
	Register()
}

type AuthHandler struct {
	svc AuthService
}

func NewAuthHandler(svc AuthService) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

type LoginResponse struct {
}

type RegisterResponse struct {
}

func (a *AuthHandler) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, LoginResponse{})
}

func (a *AuthHandler) Register(c echo.Context) error {
	return c.JSON(http.StatusOK, RegisterResponse{})
}
