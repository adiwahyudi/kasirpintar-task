package main

import (
	"fmt"
	"kasirpintar/controllers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	g := gin.Default()

	fmt.Println("Server running...")

	// Log Report
	f, err := os.OpenFile("logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("Server running...")

	g.POST("/registration", controllers.Registration)
	g.POST("/inquiry", controllers.StatusInquiry)
	g.POST("/callback", controllers.Callback)
	g.GET("/payment", controllers.Payment)
	g.Run(":8080")
}
