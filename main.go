package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-email-authentication/pkg/router"
)

func main() {
	r := gin.Default()

	router.Router(r)

	err := r.Run(":8085")
	if err != nil {
		fmt.Println("start failed!")
	}
}
