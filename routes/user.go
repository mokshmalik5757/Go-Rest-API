package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mokshmalik5757/Go-Rest-API/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controller.GetUsers)
	router.POST("/user",controller.CreateUser)
	router.DELETE("/user/:id", controller.DeleteUser)
	router.PUT("/user/:id", controller.UpdateUser)
}