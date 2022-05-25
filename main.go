package main

import (
	"github.com/gin-gonic/gin"
	m "go-email-identify/pkg/model"
	s "go-email-identify/pkg/service"
	"net/http"
)

func main() {
	r := setupRouter()
	r.Run(":8085")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/user/login", func(c *gin.Context) {
		var login m.UserDto
		if err := c.ShouldBind(&login); err == nil {
			status, resp := s.Login(login)
			if status == http.StatusOK {
				c.String(http.StatusOK, resp)
			} else {
				c.String(http.StatusInternalServerError, resp)
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.GET("/user/sendCode", func(c *gin.Context) {
		user_email := c.Query("email")
		status, resp := s.SendEmail(user_email)
		if status == http.StatusOK {
			c.String(http.StatusOK, resp)
		} else {
			c.String(http.StatusInternalServerError, resp)
		}
	})

	return router
}
