package authorize

// CreateAuthorization holds data to make an authorization request to auth0 /authorize endpoint
// ResponseType must be set to code according to RFC: https://tools.ietf.org/html/rfc6749#section-4.1.1
type CreateAuthorization struct {
	Audience     string `json:"audience"`
	ClientID     string `json:"client_id" validate:"required"`
	Nonce        string `json:"nonce"`
	RedirectURI  string `json:"redirect_uri" validate:"required"`
	ResponseType string `json:"response_type" validate:"required,eq=code"`
	Scope        string `json:"scope"`
	State        string `json:"state,omitempty"`
}
