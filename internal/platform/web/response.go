package web

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pkg/errors"
	validator "gopkg.in/go-playground/validator.v9"

	apiContext "github.com/jbelmont/api-workshop/internal/platform/context"
)

var validate = validator.New()

var (
	// ErrNotAuthorized occurs when the call is not authorized.
	ErrNotAuthorized = errors.New("Not authorized")

	// ErrDBNotConfigured occurs when the DB is not initialized.
	ErrDBNotConfigured = errors.New("DB not initialized")

	// ErrNotFound is abstracting the mgo not found error.
	ErrNotFound = errors.New("Entity not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in it's proper form")

	// ErrValidation occurs when there are validation errors.
	ErrValidation = errors.New("Validation errors occurred")
)

// Invalid describes a validation error belonging to a specific field.
type Invalid struct {
	Field string `json:"field_name"`
	Error string `json:"error"`
}

// InvalidError is a custom error type for invalid fields.
type InvalidError []Invalid

// Error implements the error interface for InvalidError.
func (err InvalidError) Error() string {
	var str string
	for _, v := range err {
		str = fmt.Sprintf("%s,{%s:%s}", str, v.Field, v.Error)
	}
	return str
}

// AppError is an application error. Its message and error code should be
// displayed to the client and served with a particular HTTP status code.
type AppError struct {
	Message string `json:"message"` // payload is in a format not supported by this method on the target resource
	Code    string `json:"code"`    // Unsupported Media Type
	Status  int    `json:"status"`  // 415
}

// Error implements the error interface
func (e AppError) Error() string {
	return e.Message
}

// JSONError is the response for errors that occur within the API.
type JSONError struct {
	Error  string       `json:"error"`
	Fields InvalidError `json:"fields,omitempty"`
}

// Unmarshal decodes the input to the struct type and checks the
// fields to verify the value is in a proper state.
func Unmarshal(r io.Reader, v interface{}) error {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		return err
	}

	var inv InvalidError
	if fve := validate.Struct(v); fve != nil {
		for _, fe := range fve.(validator.ValidationErrors) {
			inv = append(inv, Invalid{
				Field: fe.Field(),
				Error: fe.Tag(),
			})
		}
		return inv
	}

	return nil
}

// Error handles all error responses for the API.
func Error(cxt context.Context, w http.ResponseWriter, err error) {
	switch errors.Cause(err) {
	case ErrNotFound:
		RespondError(cxt, w, err, http.StatusNotFound)
		return

	case ErrInvalidID:
		RespondError(cxt, w, err, http.StatusBadRequest)
		return

	case ErrValidation:
		RespondError(cxt, w, err, http.StatusBadRequest)
		return

	case ErrNotAuthorized:
		RespondError(cxt, w, err, http.StatusUnauthorized)
		return
	}

	switch e := errors.Cause(err).(type) {
	case InvalidError:
		v := JSONError{
			Error:  "field validation failure",
			Fields: e,
		}

		Respond(cxt, w, v, http.StatusBadRequest)
		return
	}

	RespondError(cxt, w, err, http.StatusInternalServerError)
}

// RespondError sends JSON describing the error
func RespondError(ctx context.Context, w http.ResponseWriter, err error, code int) {
	Respond(ctx, w, JSONError{Error: err.Error()}, code)
}

// Respond sends JSON to the client.
// If code is StatusNoContent, v is expected to be nil.
func Respond(ctx context.Context, w http.ResponseWriter, data interface{}, code int) {

	// Set the status code for the request logger middleware.
	v := ctx.Value(apiContext.KeyValues).(*apiContext.Values)
	v.StatusCode = code

	// Just set the status code and we are done.
	if code == http.StatusNoContent {
		w.WriteHeader(code)
		return
	}

	// Set the content type.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response and context.
	w.WriteHeader(code)

	// Marshal the data into a JSON string.
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("%s : Respond %v Marshalling JSON response\n", v.TraceID, err)
		jsonData = []byte("{}")
	}

	// Send the result back to the client.
	io.WriteString(w, string(jsonData)) // nolint: errcheck
}
