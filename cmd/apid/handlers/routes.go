package handlers

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/jbelmont/api-workshop/cmd/apid/config"
	"github.com/jbelmont/api-workshop/internal/platform/database"
)

// API returns a handler for a set of routes.
func API(masterDB *database.DB, cfg config.Config) http.Handler {
	// Create the application.
	router := httptreemux.New()
	api := router.NewGroup("/api/v1")

	// Instantiate the handlers
	h := Hero{
		config:   cfg,
		MasterDB: masterDB,
	}

	api.GET("/heroes", h.List)
	api.POST("/heroes", h.Create)
	api.GET("/heroes/:id", h.Retrieve)
	api.PUT("/heroes/:id", h.Update)
	api.DELETE("/heroes/:id", h.Delete)

	return router
}
