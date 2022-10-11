package models

import (
	"context"
	"fmt"
	"schedule-events/packages/common/services"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventSchedule struct {
	ID          primitive.ObjectID     `bson:"_id"`
	Name        string                 `json:"name" binding:"required"`
	Description string                 `json:"description" binding:"required"`
	Origin      string                 `json:"origin" binding:"required"`
	Date        primitive.DateTime     `primitive:"date" binding:"required"`
	Data        map[string]interface{} `bson:"data"  binding:"required"`
}

type dbInsertion struct {
	Status bool
}

type dbOperation struct {
	Status bool
}

var eventScheduleCollection *mongo.Collection = services.GetCollection(services.DB, "eventSchedules")

//From Http

func InsertEventSchedule(c *gin.Context, newEventSchedule EventSchedule) dbInsertion {
	result, err := eventScheduleCollection.InsertOne(c, newEventSchedule)

	if err != nil {
		fmt.Println("Error inserting on db", err)
		return dbInsertion{Status: false}
	}

	fmt.Println("Insertion successfully", result)
	return dbInsertion{Status: true}
}

//From cron job

func LocalGetEventSchedules() (eventSchedules []EventSchedule, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	results, err := eventScheduleCollection.Find(ctx, bson.M{"date": bson.M{
		"$lte": primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0)),
	}})

	if err != nil {
		return nil, err
	}
	err = results.All(context.TODO(), &eventSchedules)
	if err != nil {
		return nil, err
	}
	return
}

func LocalDeleteEventSchedule(eventSchedule EventSchedule) dbOperation {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	eventId := eventSchedule.ID

	result, err := eventScheduleCollection.DeleteOne(ctx, bson.M{"_id": eventId})

	if err != nil {
		fmt.Println("Error deleting on db", err)
		return dbOperation{Status: false}
	}

	fmt.Println("Delete successfully", result)
	return dbOperation{Status: true}
}
