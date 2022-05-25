package service

import (
	"github.com/gin-gonic/gin"
	email_utils "go-email-identify/pkg/utils"
	"net/http"
)

func SendEmail(c *gin.Context) (int, string) {
	email_utils.SendMessage()
	return http.StatusOK, "send email success"
}
