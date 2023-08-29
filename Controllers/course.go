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

// Create Course Data
func CreateCourse(c echo.Context) error {
    db := Config.GetDB()
    
    course := new(Models.Course)

    if err := c.Bind(course); err != nil {
        data := map[string]interface{}{
            "message": err.Error(),
        }
        return c.JSON(http.StatusInternalServerError, data)
    }

    if err := db.Create(&course).Error; err != nil {
        data := map[string]interface{}{
            "message": err.Error(),
        }
        return c.JSON(http.StatusInternalServerError, data)
    }

    response := map[string]interface{}{
        "data": course,
    }
    return c.JSON(http.StatusOK, response)
}