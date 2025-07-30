package main

import (
	"task-manager-ca/Delivery/routers"
	"task-manager-ca/infrastructure"

	"github.com/gin-gonic/gin"
	
)

func main() {

    // connect database
    prodDB := infrastructure.DbInit()

    // create a gin.Engine instance
    router := gin.Default()

    // Start your HTTP server
    routers.SetUp(prodDB, router)
    router.Run("localhost:8000")
   
}
