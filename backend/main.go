package main

import (
	"go-url-shortener/config"
	"go-url-shortener/models"
	"go-url-shortener/routes"

	"github.com/gin-gonic/gin"
)

func main(){
    //1. Connect the DB
    config.ConnectDB()

    config.DB.AutoMigrate(&models.Link{})
    
    //2. Init GIN
    r:=gin.Default()
    //3. Setup Routes Function
    routes.SetupRoutes(r)
    //4. Start the Server
    r.Run(":8000")
}