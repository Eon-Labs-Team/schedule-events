package main

import (
	"schedule-events/packages/common/services"
	"schedule-events/packages/cronJobs"
	"schedule-events/packages/eventSchedule/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	services.ConnectDB()
	cronJobs.StartCron()
	router := gin.Default()

	//routes insert schedule-event
	routes.RoutesConfig(router)

	router.Run("localhost:8080")
}
