package main

import (
	"backend/config"
	"backend/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dbConfig := config.GetDBConfig()
	httpConfig := config.GetHTTPConfig()
	dbConn := db.Connect(dbConfig)
	if err := db.Migrate(dbConn); err != nil {
		panic("Database migration failed: " + err.Error())
	}

	e := echo.New()

	e.Use(middleware.Logger())

	// routes.RegisterRoutes(e, dbConn)

	serverAddress := httpConfig.Host + ":" + httpConfig.Port
	e.Logger.Fatal(e.Start(serverAddress))
}
