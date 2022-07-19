package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/login-service/entity"
)

//get user data from database
func GetUser(ctx *gin.Context) {

	var p entity.UserProfile

	fmt.Println(p)
	ctx.JSON(200, p)
}
func SendOTP(ctx *gin.Context) {
	//use twillio here to send otp

}

func ValidateOTP(c *gin.Context) {
	//verify otp received from user
	//update isVerified cloumn to true
	var otp entity.OTP
	if err := c.ShouldBindJSON(&otp); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, "ok")
}
func CreateUser(c *gin.Context) {

	var u entity.UserProfile
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, u)
}
func Login(ctx *gin.Context) {
	//perform login process here
	//1. verify if the phone number/ user exists in the db
	//2. if exists, send otp

}