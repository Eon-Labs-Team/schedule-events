package routes

import (
	"schedule-events/packages/eventSchedule/controllers"

	"github.com/gin-gonic/gin"
)

func RoutesConfig(router *gin.Engine) {
	router.POST("/eventSchedule", controllers.InsertEventSchedule)
}
