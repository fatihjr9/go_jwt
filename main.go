package main

import (
	"jwt_go/controller"
	"jwt_go/middleware"
	"jwt_go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDb()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/login", controller.Login)
	public.POST("/register", controller.Register)

	protected := r.Group("/api/admin")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/user", controller.CurrentUser)

	r.Run(":8080")

}
