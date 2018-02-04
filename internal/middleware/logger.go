package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	apiContext "github.com/jbelmont/api-workshop/internal/platform/context"
	"github.com/jbelmont/api-workshop/internal/platform/web"
)

// RequestLogger writes some information about the request to the logs in
// the format: TraceID : (200) GET /foo -> IP ADDR (latency)
func RequestLogger(next web.Handler) web.Handler {

	// Wrap this handler around the next one provided.
	h := func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
		v := ctx.Value(apiContext.KeyValues).(*apiContext.Values)

		err := next(ctx, w, r, params)

		log.Printf("%s : (%d) : %s %s -> %s (%s)",
			v.TraceID,
			v.StatusCode,
			r.Method, r.URL.Path,
			r.RemoteAddr, time.Since(v.Now),
		)

		return err
	}

	return h
}
