package service

import (
	"go-email-identify/pkg/model"
	"net/http"
)

func Login(dto model.UserDto) (int, string) {
	return http.StatusOK, dto.User
}
