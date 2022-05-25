package service

import (
	email_utils "go-email-identify/pkg/utils"
	"net/http"
)

func SendEmail(email string) (int, string) {
	email_utils.SendMessage(email)
	return http.StatusOK, "send email success"
}
