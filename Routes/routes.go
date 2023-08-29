package Routes

import (
	"simple-rest-go-echo/Controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/courses", Controllers.GetCourse)
}