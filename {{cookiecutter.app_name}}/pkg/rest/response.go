package rest

import (
	"errors"
	"net/http"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/errr"
)

type Envelope struct {
	Success bool
	Message string
	Data    interface{}     `json:",omitempty"`
	Paging  *ResponsePaging `json:",omitempty"`
}

func ParseErrHttp(err error) (status int, msg string) {

	errorMessage := errr.GetLastNErrorMessage(err, 2)

	var stcErr errr.StatusCodeError
	if errors.As(err, &stcErr) {
		return stcErr.StatusCode, errorMessage
	}
	if errors.Is(err, errr.ErrNoRows) {
		return http.StatusBadRequest, errorMessage
	}
	return 500, errorMessage
}

/* func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message string) {
	// traceID := ReadTraceID(r.Context())

	env := Envelope{
		message: message,
		data:    nil,
	}
	err := WriteJSON(w, status, env, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ErrorPayloadResponse(w http.ResponseWriter, r *http.Request, message string) {
	// traceID := ReadTraceID(r.Context())

	env := Envelope{
		err: message,
		data: nil,
	}
	err := WriteJSON(w, http.StatusBadRequest, env, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
} */
