package main

import (
	"github.com/yaseer08/Go_Practice/User_CRUD/controllers"
	"github.com/yaseer08/Go_Practice/User_CRUD/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.GET("/users", controllers.GetAllUsers)

	// Add other routes...

	r.Run(":8080")
}
