package main

import (
	"login-service/database"
	"login-service/router"

	_ "github.com/lib/pq"
)

//var db *sql.DB

func main() {
	//TODO :
	// Create an api /user GET encrypted with JWT for login
	database.CreatConnectionTodb()
	router.InitRouter()
}
