package web

import (
	"context"
	"net/http"
	"time"

	"github.com/dimfeld/httptreemux"
	"github.com/pborman/uuid"

	apiContext "github.com/jbelmont/api-workshop/internal/platform/context"
)

// A Handler is a type that handles an http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct
type App struct {
	*httptreemux.TreeMux
	mw []Middleware
}

// New creates an App value that handle a set of routes for the application.
// You can provide any number of middleware and they'll be used to wrap every
// request handler.
func New(mw ...Middleware) *App {
	return &App{
		TreeMux: httptreemux.New(),
		mw:      mw,
	}
}

// Use adds the set of provided middleware onto the Application middleware
// chain. Any route running off of this App will use all the middleware provided
// this way always regardless of the ordering of the Handle/Use functions.
func (a *App) Use(mw ...Middleware) {
	a.mw = append(a.mw, mw...)
}

// Handle is our mechanism for mounting Handlers for a given HTTP verb and path
// pair, this makes for really easy, convenient routing.
func (a *App) Handle(verb, path string, handler Handler, mw ...Middleware) {

	// Wrap up the application-wide first, this will call the first function
	// of each middleware which will return a function of type Handler. Each
	// Handler will then be wrapped up with the other handlers from the chain.
	handler = wrapMiddleware(wrapMiddleware(handler, mw), a.mw)

	// The function to execute for each request.
	h := func(w http.ResponseWriter, r *http.Request, params map[string]string) {

		// Set the context with the required values to
		// process the request.
		v := apiContext.Values{
			TraceID: uuid.New(),
			Now:     time.Now(),
		}
		ctx := context.WithValue(r.Context(), apiContext.KeyValues, &v)

		// Set the trace id on the outgoing requests before any other header to
		// ensure that the trace id is ALWAYS added to the request regardless of
		// any error occurring or not.
		w.Header().Set(apiContext.TraceIDHeader, v.TraceID)

		// Call the wrapped handler functions.
		handler(ctx, w, r, params) // nolint: errcheck
	}

	// Add this handler for the specified verb and route.
	a.TreeMux.Handle(verb, path, h)
}
