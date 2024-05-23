package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Rafaeldias0934/primeiro-crud-go.git/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("minha aplicação é" + os.Getenv("TESTE"))

	router := gin.Default()

	routes.InirRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
