package Routes

import (
	"simple-rest-go-echo/Controllers"

	"github.com/labstack/echo/v4"
)

// setup for routing
func SetupRoutes(e *echo.Echo) {
	e.GET("/courses", Controllers.GetCourse)
	e.POST("/courses", Controllers.CreateCourse)
	e.PUT("/courses/:id", Controllers.UpdateCourse)
	e.DELETE("/courses/:id", Controllers.DeleteCourse)
}