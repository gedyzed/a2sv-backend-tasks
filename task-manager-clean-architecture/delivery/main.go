package main

import (
	"task-manager-ca/Delivery/routers"
	"task-manager-ca/infrastructure"

	"github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
	
)

func main() {

    if err := godotenv.Load(); err != nil {
        log.Println(" No .env file found or failed to load it")
    }

    // connect database
    prodDB := infrastructure.DbInit()

    // create a gin.Engine instance
    router := gin.Default()

    // Start your HTTP server
    routers.SetUp(prodDB, router)
    router.Run("localhost:8000")
   
}
