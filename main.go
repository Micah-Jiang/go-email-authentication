package main

import (
	"github.com/gin-gonic/gin"
	m "go-email-authentication/pkg/model"
	s "go-email-authentication/pkg/service"
	"net/http"
)

func main() {
	r := setupRouter()
	r.Run(":8085")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/user/register", func(c *gin.Context) {
		var user m.UserDto
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&user); err == nil {
			s.Register(user)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

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
		username := c.Query("username")
		status, resp := s.SendEmail(username, user_email)
		if status == http.StatusOK {
			c.String(http.StatusOK, resp)
		} else {
			c.String(http.StatusInternalServerError, resp)
		}
	})

	return router
}
