package main

import (
	_ "github.com/lib/pq"
	"github.com/login-service/database"
	"github.com/login-service/router"
)

//var db *sql.DB

func main() {
	//TODO :
	// Create an api /user GET encrypted with JWT for login
	database.CreatConnectionTodb()
	router.InitRouter()
}
