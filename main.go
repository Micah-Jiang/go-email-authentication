package main

import (
	"github.com/gin-gonic/gin"
	m "go-email-authentication/pkg/dto"
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
		var userDto m.UserDto
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&userDto); err == nil {
			statusCode, msg := s.Register(userDto)
			c.JSON(http.StatusOK, gin.H{
				"code": statusCode,
				"msg":  msg,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.GET("/user/login", func(c *gin.Context) {
		var userDto m.UserDto
		if err := c.ShouldBind(&userDto); err == nil {
			status, resp := s.Login(userDto, c)
			c.JSON(http.StatusOK, gin.H{
				"code": status,
				"msg":  resp,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.GET("/user/sendCode", func(c *gin.Context) {
		email := c.Query("email")
		username := c.Query("username")
		status, resp := s.SendEmail(username, email)
		if status == http.StatusOK {
			c.String(http.StatusOK, resp)
		} else {
			c.String(http.StatusInternalServerError, resp)
		}
	})

	return router
}
