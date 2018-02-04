package context

import "time"

// TraceIDHeader is the header added to outgoing requests which adds the traceID to it.
const TraceIDHeader = "X-Workshop-Trace-ID"

// Key represents the type of value for the context key.
type ctxKey int

// KeyValues is how request values or stored/retrieved.
const KeyValues ctxKey = 1

// Values hold the state for each request.
// As a contract, the UserID must be an ObjectIdHex. This is to keep backward compatibility (mongoDB ObjectId)
//  and avoid the web package depending on bson.
type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
	ID         string
	Name       string
}
