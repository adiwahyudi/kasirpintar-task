package main

import (
	"fmt"
	"kasirpintar/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	g := gin.Default()

	fmt.Println("Gin Success!")

	g.POST("/registration", controllers.Registration)

	g.Run(":8080")
}
