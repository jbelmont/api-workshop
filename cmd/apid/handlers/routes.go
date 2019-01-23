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

	// Regular CRUD Endpoints for Hero Resources
	api.GET("/heroes", h.List)
	api.POST("/heroes", h.Create)
	api.GET("/heroes/:id", h.Retrieve)
	api.PUT("/heroes/:id", h.Update)
	api.DELETE("/heroes/:id", h.Delete)

	// SSO / OpenID Connect with OAuth2 using Auth0
	// These endpoints are in support of Authorization Code Flow
	// https://openid.net/specs/openid-connect-core-1_0.html#AuthorizationEndpoint

	auth := Authorize{
		config: cfg,
	}
	api.GET("/authorize", auth.Authorize)

	return router
}
