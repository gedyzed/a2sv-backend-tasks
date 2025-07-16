package main
import "task_manager/router"

func main(){

	//start the server
	router.Routers().Run("localhost:8000")
}