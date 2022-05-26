package service

import (
	"go-email-authentication/pkg/utils"
	"net/http"
)

func SendEmail(username, email string) (int, string) {
	err := utils.SendMessage(username, email)
	if err != nil {
		return http.StatusInternalServerError, "send email failed"
	}
	return http.StatusOK, "send email success"
}
