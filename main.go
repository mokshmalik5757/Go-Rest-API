package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mokshmalik5757/Go-Rest-API/config"
	"github.com/mokshmalik5757/Go-Rest-API/routes"
	"github.com/mokshmalik5757/Go-Rest-API/server"
)


func main() {
	router:= gin.Default()
	config.Connect() //connecting with the database
	server.HealthCheckRoute(router)
	routes.UserRoute(router)
	router.Run("localhost:8080")
}