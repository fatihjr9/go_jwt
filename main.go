package main

import (
	"jwt_go/controller"
	"jwt_go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDb()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controller.Register)

	r.Run(":8080")

}
