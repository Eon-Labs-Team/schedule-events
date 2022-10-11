package controllers

import (
	"net/http"
	"schedule-events/packages/eventSchedule/models"
	"schedule-events/packages/eventSchedule/responses"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert a EventSchedule into mongoDb database
func InsertEventSchedule(ctx *gin.Context) {
	var newEventSchedule models.EventSchedule
	newEventSchedule.ID = primitive.NewObjectID()

	// Call BindJSON to bind the received JSON to
	if err := ctx.BindJSON(&newEventSchedule); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.EventScheduleResponse{Status: http.StatusBadRequest, Message: "Wrong event data"})
		return
	}

	// Add the new album to the slice.
	var result = models.InsertEventSchedule(ctx, newEventSchedule)
	if !result.Status {
		ctx.JSON(http.StatusBadRequest, responses.EventScheduleResponse{Status: http.StatusBadRequest, Message: "Error inserting event"})
		return
	}

	ctx.JSON(http.StatusCreated, responses.EventScheduleResponse{Status: http.StatusCreated, Message: "Event scheduled correctly"})
	return
}
