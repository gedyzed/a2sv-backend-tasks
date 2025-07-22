package main

import (
	"task_manager_with_auth/data"
	"task_manager_with_auth/router"
)

func main() {

	data.DatabaseConnection()
	router.Routers().Run("localhost:8080")
}