package tdtg

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func addAPI(g *echo.Group) error {
	return nil
}

func createApp() (*echo.Echo, error) {
	// Create the echo application...
	app := echo.New()
	app.HideBanner = true
	app.HidePort = true

	// Add middleware...
	app.Use(middleware.Logger())  // TODO: Switch to custom logger?
	app.Use(middleware.Recover()) // TODO: Check for running in prod?
	app.Use(middleware.RemoveTrailingSlash())

	// Serve static files...

	// Add the endpoints...
	g := app.Group("/api")
	if err := addAPI(g); err != nil {
		return nil, err
	}

	// Return the application...
	return app, nil
}
