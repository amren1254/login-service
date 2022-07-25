package router

import (
	"login-service/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.GET("/user/:phonenumber", controller.GetUser)
	router.POST("/sendotp/:phonenumber", controller.SendOTP)
	router.POST("validateotp", controller.ValidateOTP)
	router.POST("user", controller.CreateUser)
	router.POST("login", controller.Login)
	router.POST("otp", controller.CreateOtp)

	router.Run()

}
