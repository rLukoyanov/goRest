package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.com/main/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	// context.Request.Body
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
	}

	event.Id = 1
	event.UserId = 1

	context.JSON(http.StatusCreated, gin.H{
		"message": "event created",
		"event":   event,
	})
}
