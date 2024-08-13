package main

import (
	"full-crud/db" // Assuming "db" is the package name for your database operations
	"full-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.InitDB() // Initialize the database connection
	routes.RegisterRoutes(server)
	// port:= 8000

	server.Run(":8080") //localhost
}
