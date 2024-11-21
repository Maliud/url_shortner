package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Maliud/url_shortner/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()

	setupRouters(router)

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run(":" + port))
}

func setupRouters(router *gin.Engine) {
	router.POST(os.Getenv("ROUTER_V1"), routes.ShortenURL)
	router.GET(os.Getenv("ROUTER_GETBYSHORTID"), routes.GetByShortID)
	router.DELETE(os.Getenv("ROUTER_GETBYSHORTID"), routes.DeleteURL)
	router.PUT(os.Getenv("ROUTER_GETBYSHORTID"), routes.EditURl)
	router.POST(os.Getenv("ROUTER_ADDTAG"), routes.AddTag)
}
