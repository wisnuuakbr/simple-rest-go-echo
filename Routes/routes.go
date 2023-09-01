package Routes

import (
	"simple-rest-go-echo/Controllers"

	"github.com/labstack/echo/v4"
)

// setup for routing
func SetupRoutes(e *echo.Echo) {
	e.GET("/course", Controllers.GetCourse)
	e.POST("/course", Controllers.CreateCourse)
	e.PUT("/course/:id", Controllers.UpdateCourse)
	e.DELETE("/course/:id", Controllers.DeleteCourse)
}