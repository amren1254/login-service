package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"login-service/database"
	"login-service/entity"
	"login-service/twillio"
)

//get user data from database
func GetUser(ctx *gin.Context) {
	phoneNumber := ctx.Param("phonenumber")
	fmt.Println(phoneNumber)
	userProfile := database.GetUser(phoneNumber)
	ctx.JSON(200, userProfile)
}
func SendOTP(ctx *gin.Context) {
	//use twillio here to send otp
	phoneNumber := ctx.Param("phonenumber")
	isInserted := database.AddPhoneNumber(phoneNumber)
	ctx.JSON(200, isInserted)
}

func ValidateOTP(c *gin.Context) {
	//verify otp received from user
	//update isVerified cloumn to true
	var otp entity.OTP
	if err := c.ShouldBindJSON(&otp); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	isValid := database.ValidateOTP(otp)

	c.JSON(200, isValid)
}
func CreateUser(c *gin.Context) {

	var u entity.UserProfile
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	database.CreateUser(u)
	c.JSON(200, u)
}
func Login(ctx *gin.Context) {
	//perform login process here
	//1. verify if the phone number/ user exists in the db
	//2. if exists, send otp

}

func CreateOtp(c *gin.Context) {
	var t entity.SendOTP
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	twillio.InitTwillio(t)

	c.JSON(200, t)
}
