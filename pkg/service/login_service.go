package service

import (
	"go-email-authentication/pkg/model"
	"net/http"
)

func Login(dto model.UserDto) (int, string) {
	return http.StatusOK, dto.User
}
