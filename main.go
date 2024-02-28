package main

import (
	"projects/golang-jwt/controllers"
	"projects/golang-jwt/initializers"
	"projects/golang-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	if err := r.Run(); err != nil {
		panic(err)
	}
}