package api

import (
	"github.com/gin-gonic/gin"
	m "go-email-authentication/pkg/dto"
	s "go-email-authentication/pkg/service"
	"net/http"
)

func Register(c *gin.Context) {
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
}

func Login(c *gin.Context) {
	var userDto m.UserDto
	if err := c.ShouldBind(&userDto); err == nil {
		s.Login(userDto, c)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
