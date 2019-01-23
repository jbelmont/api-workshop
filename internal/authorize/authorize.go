package authorize

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/jbelmont/api-workshop/cmd/apid/config"
)

const (
	scope = "profile email"
)

// Authorize will help obtain consent from a user to invoke the API (specified in audience)
// and execute certain actions (specified in scope) on behalf of the user.
func Authorize(cAuth CreateAuthorization, cfg config.Config) (*string, error) {
	cAuth.RedirectURI = cfg.Auth0CallbackURL + "/" + url.PathEscape(cAuth.RedirectURI)
	cAuth.Audience = cfg.Auth0Audience
	cAuth.Scope = scope

	query := prepareQueryStringParametersForAuthorizeQuery(cAuth)
	request, err := http.NewRequest("GET", cfg.Auth0URL+"/authorize"+query, nil)
	if err != nil {
		return nil, errors.Wrap(err, "http request to authorize endpoint")
	}
	url := request.URL.String()
	return &url, nil
}

// creates a query string parameter for the GET /authorize endpoint and escape parts of query
func prepareQueryStringParametersForAuthorizeQuery(createAuthorization CreateAuthorization) string {
	return "?" + "audience=" + url.QueryEscape(createAuthorization.Audience) +
		"&scope=" + url.QueryEscape(createAuthorization.Scope) +
		"&response_type=" + url.QueryEscape(createAuthorization.ResponseType) +
		"&client_id=" + url.QueryEscape(createAuthorization.ClientID) +
		"&redirect_uri=" + url.QueryEscape(createAuthorization.RedirectURI) +
		"&state=" + url.QueryEscape(createAuthorization.State)
}
