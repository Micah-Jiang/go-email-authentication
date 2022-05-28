package router

import (
	"github.com/gin-gonic/gin"
	"go-email-authentication/pkg/api"
)

func Router(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/login", api.Login)
		user.GET("/register", api.Register)
	}

	sms := r.Group("/sms")
	{
		sms.GET("/sendCode", api.SendSmsCode)
	}
}
