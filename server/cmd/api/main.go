package main

import (
	"fmt"
	"money-backend/internal/routes"
	"money-backend/pkg/database"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

func main() {
	database.InitDB()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Origin Nuxt Anda
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/", helloHandler)
	routes.SetupRoute(r)

	port := ":8000"
	fmt.Println("Server berjalan di http://localhost" + port)

	r.Run(port)
}
