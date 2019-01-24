package handlers

import (
	"log"
	"net/http"

	"github.com/jbelmont/api-workshop/cmd/apid/config"
	"github.com/jbelmont/api-workshop/internal/authorize"
)

// Authorize holds information for an authorization flow in OAuth2/OpenID Connect
type Authorize struct {
	config config.Config
}

// Authorize is a handler method that is used for the /authorize endpoint in auth0.
func (a *Authorize) Authorize(w http.ResponseWriter, r *http.Request, params map[string]string) {
	queryStringParameters := r.URL.Query()
	createAuthorization := authorize.CreateAuthorization{
		ClientID:     queryStringParameters.Get("client_id"),
		Nonce:        queryStringParameters.Get("nonce"),
		ResponseType: queryStringParameters.Get("response_type"),
		State:        queryStringParameters.Get("state"),
	}

	url, err := authorize.Authorize(createAuthorization, a.config)
	if err != nil {
		log.Printf("%s: Calling Authorize Service Function", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, *url, http.StatusFound)
	return
}
