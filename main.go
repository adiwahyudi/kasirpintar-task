package main

import (
	"kasirpintar/controllers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Logging
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	g := gin.Default()
	log.Println("Server running...")

	g.POST("/registration", controllers.Registration)
	g.POST("/inquiry", controllers.StatusInquiry)
	g.POST("/payment", controllers.Payment)
	// g.POST("/callback", controllers.Callback)
	g.Run(":8080")
}
