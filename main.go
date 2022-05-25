package main

import (
	"github.com/gin-gonic/gin"
	email_service "go-email-identify/pkg/service"
	"net/http"
)

func main() {
	r := setupRouter()
	r.Run(":8085")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/login", func(c *gin.Context) {
		status, resp := email_service.SendEmail(c)

		if status == http.StatusOK {
			c.String(http.StatusOK, resp)
		} else {
			c.String(http.StatusInternalServerError, resp)
		}
	})

	return router
}
