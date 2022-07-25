package main

import (
	"fmt"
	"login-service/database"
	"login-service/router"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//var db *sql.DB

func main() {
	//TODO :
	// Create an api /user GET encrypted with JWT for login
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")

	}
	database.CreatConnectionTodb()
	router.InitRouter()
}
