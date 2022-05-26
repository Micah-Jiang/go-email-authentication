package service

import (
	"go-email-authentication/pkg/model"
	"net/http"
)

func Login(dto model.UserDto) (int, string) {
	// 1. check username and password

	// 2. return ret
	return http.StatusOK, dto.User
}

func Register(dto model.UserDto) (int, string) {
	// 1. check username and password

	// 2. return ret
	return http.StatusOK, dto.User
}
