package main

import (
	"task_manager_with_mongodb/router"
	"task_manager_with_mongodb/data"

)


func main(){

	// mongodb connection
	data.DatabaseConnection()

	//start the server
	router.Routers().Run("localhost:8000")
	
}


