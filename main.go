package main

import (
	"github.com/3AM-Developer/dae/controllers"
	"github.com/3AM-Developer/dae/database"
	"github.com/3AM-Developer/dae/initializers"
	"github.com/3AM-Developer/dae/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	database.ConnectToDb()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.POST("/signup", middleware.CheckPreSignedURL, controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
