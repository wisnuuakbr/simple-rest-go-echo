package main

import (
	"simple-rest-go-echo/Config"
	"simple-rest-go-echo/Models"
	"simple-rest-go-echo/Routes"

	"github.com/labstack/echo/v4"
)

func main() {
    // Initialize Echo instance
    e := echo.New()

    /// Connect to the database
    Config.DatabaseInit()
    defer Config.GetDB().DB()

    // Perform migrations using AutoMigrate
    db := Config.GetDB()
    err := db.AutoMigrate(&Models.Course{})
    if err != nil {
        panic(err)
    }
    
    // Set up Routes
    Routes.SetupRoutes(e)

    // Start the server
    e.Logger.Fatal(e.Start(":8080"))
}