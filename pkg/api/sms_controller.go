package api

import (
	"github.com/gin-gonic/gin"
	s "go-email-authentication/pkg/service"
	"net/http"
)

func SendSmsCode(c *gin.Context) {
	email := c.Query("email")
	username := c.Query("username")
	status, resp := s.SendEmail(username, email)
	if status == http.StatusOK {
		c.String(http.StatusOK, resp)
	} else {
		c.String(http.StatusInternalServerError, resp)
	}
}
