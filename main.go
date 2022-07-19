package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "userlogin"
)

type UserProfile struct {
	FullName    string `json:"fullname"`
	EmailId     string `json:"emailid"`
	PhoneNumber string `json:"phonenumber"`
}

var db *sql.DB

func createConnection() {
	//get user data from database
	//connect with database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	// err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected successfully")
}
func getUser(ctx *gin.Context) {

	var p UserProfile
	db.QueryRow("SELECT fullname, emailid, phonenumber FROM userprofile where phonenumber='88967264840'").Scan(&p.FullName, &p.EmailId, &p.PhoneNumber)
	fmt.Println(p)
	ctx.JSON(200, p)
}
func sendOTP(ctx *gin.Context) {
	//use twillio here to send otp

}

type OTP struct {
	PhoneNumber string `json:"phonenumber"`
	Otp         string `json:"otp"`
}

func validateOTP(c *gin.Context) {
	//verify otp received from user
	//update isVerified cloumn to true
	var otp OTP
	if err := c.ShouldBindJSON(&otp); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE userprofile SET isverified=$1 WHERE phonenumber=$2",
		true, otp.PhoneNumber)
	if err != nil {
		log.Println("Db Update failed", err.Error())
	}
	c.JSON(200, "ok")
}
func createUser(c *gin.Context) {
	var u UserProfile
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	var isverified bool = false
	db.QueryRow("SELECT isverified FROM userprofile where phonenumber=$1", u.PhoneNumber).Scan(&isverified)
	if isverified {
		_, err := db.Exec("UPDATE userprofile SET fullname=$1, emailid=$2 WHERE phonenumber=$3",
			u.FullName, u.EmailId, u.PhoneNumber)
		if err != nil {
			log.Println("Db Update failed", err.Error())
		}
	}

	//check if isVerified column is true, create user else don't

	c.JSON(200, u)
}
func login(ctx *gin.Context) {
	//perform login process here
	//1. verify if the phone number/ user exists in the db
	//2. if exists, send otp

}
func main() {
	//TODO :
	// Create an api /user GET encrypted with JWT for login
	createConnection()
	router := gin.Default()
	router.GET("/user", getUser)
	router.POST("/sendotp", sendOTP)
	router.POST("validateotp", validateOTP)
	router.POST("user", createUser)
	router.POST("login", login)

	router.Run()

}
