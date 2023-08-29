package Controllers

import (
	"net/http"
	"simple-rest-go-echo/Config"
	"simple-rest-go-echo/Models"

	"github.com/labstack/echo/v4"
)

// Get All Courses Data
func GetCourse(c echo.Context) error {
    // Get the database instance from Config package
    db := Config.GetDB()

    // define slice
    var course []*Models.Course

    if res := db.Find(&course); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}
		return c.JSON(http.StatusOK, data)
	}

    response := map[string]interface{}{
		"data": course,
	}

    return c.JSON(http.StatusOK, response)

}