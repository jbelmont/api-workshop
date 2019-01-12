package handlers

import (
	"net/http"

	"github.com/jbelmont/api-workshop/cmd/apid/config"
	"github.com/jbelmont/api-workshop/internal/middleware"
	database "github.com/jbelmont/api-workshop/internal/platform/db"
	"github.com/jbelmont/api-workshop/internal/platform/web"
)

// API returns a handler for a set of routes.
func API(masterDB *database.DB, cfg config.Config) http.Handler {
	// Create the application.
	app := web.New(middleware.RequestLogger)

	h := Hero{
		MasterDB: masterDB,
	}

	app.Handle("POST", "/api/v1/heroes", h.Create)
	app.Handle("GET", "/api/v1/heroes", h.List)
	app.Handle("GET", "/api/v1/heroes/:id", h.Retrieve)
	app.Handle("DELETE", "/api/v1/heroes/:id", h.Delete)

	return app
}
