package controllers

import (
	"log"
	"net/http"

	"github.com/yaseer08/Go_Practice/User_CRUD/db"
	"github.com/yaseer08/Go_Practice/User_CRUD/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = db.Client.Database("testdb").Collection("users")

func GetAllUsers(c *gin.Context) {
	var users []models.User
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		log.Printf("Error while getting all users: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal server error"})
		return
	}

	// Close the cursor once finished
	defer cursor.Close(c)

	for cursor.Next(c) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// Add other CRUD functions (CreateUser, GetUserByID, UpdateUser, DeleteUser) similarly...
