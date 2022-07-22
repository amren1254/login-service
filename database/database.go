package database

import (
	"database/sql"
	"fmt"
	"log"

	"login-service/entity"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "userlogin"
)

var db *sql.DB

func CreatConnectionTodb() {

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
func AddPhoneNumber(phoneNumber string) bool {
	_, err := db.Exec("INSERT INTO userprofile (phonenumber) VALUES ($1)",
		phoneNumber)
	if err != nil {
		log.Println("Phone Number insertion to db failed", err.Error())
		return false
	}
	return true
}

//Get user from db
func GetUser(phoneNumber string) entity.UserProfile {
	var p entity.UserProfile
	db.QueryRow("SELECT fullname, emailid, phonenumber FROM userprofile where phonenumber=$1", phoneNumber).Scan(&p.FullName, &p.EmailId, &p.PhoneNumber)
	return p
}

func ValidateOTP(otp entity.OTP) bool {
	_, err := db.Exec("UPDATE userprofile SET isverified=$1 WHERE phonenumber=$2",
		true, otp.PhoneNumber)
	if err != nil {
		log.Println("Db Update failed", err.Error())
		return false
	}
	return true
}

//check if isVerified column is true, create user else don't
func CreateUser(u entity.UserProfile) {
	var isverified bool = false
	db.QueryRow("SELECT isverified FROM userprofile where phonenumber=$1", u.PhoneNumber).Scan(&isverified)
	if isverified {
		_, err := db.Exec("UPDATE userprofile SET fullname=$1, emailid=$2 WHERE phonenumber=$3",
			u.FullName, u.EmailId, u.PhoneNumber)
		if err != nil {
			log.Println("Db Update failed", err.Error())
		}
	}
}
