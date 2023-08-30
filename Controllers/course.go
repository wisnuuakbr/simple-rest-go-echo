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

// Update Courses Data
func UpdateCourse(c echo.Context) error {
    db := Config.GetDB()

    course := new(Models.Course)

    // Binding data
    if err := c.Bind(course); err != nil {
        data := map[string]interface{}{
            "message": err.Error(),
        }
        return c.JSON(http.StatusInternalServerError, data)
    }
    
    // Retrieve ID from the URL parameter
    id := c.Param("id")
    existing_course := new(Models.Course)
    
    if err := db.First(&existing_course, id).Error; err != nil {
        data := map[string]interface{}{
            "message": err.Error(),
        }
        return c.JSON(http.StatusInternalServerError, data)
    }

    existing_course.Name = course.Name
	existing_course.Description = course.Description
	existing_course.Price = course.Price

    if err := db.Save(&existing_course).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

    response := map[string]interface{}{
        "data": existing_course,
    }
    return c.JSON(http.StatusOK, response)
}

// Delete Course Data
func DeleteCourse(c echo.Context) error {
    db := Config.GetDB()

    course := new(Models.Course)

    id := c.Param("id")

    err := db.Delete(&course, id).Error
    if err != nil {
        data := map[string]interface{}{
            "message": err.Error(),
        }
        return c.JSON(http.StatusInternalServerError, data)
    }

    response := map[string]interface{}{
        "message": "Data berhasil dihapus!",
    }
    return c.JSON(http.StatusOK, response)
}