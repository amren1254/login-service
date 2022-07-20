package router

import (
	"github.com/gin-gonic/gin"
	"github.com/login-service/controller"
)

func InitRouter() {
	router := gin.Default()
	router.GET("/user", controller.GetUser)
	router.POST("/sendotp", controller.SendOTP)
	router.POST("validateotp", controller.ValidateOTP)
	router.POST("user", controller.CreateUser)
	router.POST("login", controller.Login)

	router.Run()

}
