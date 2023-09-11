package xrender

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/xerrors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.opentelemetry.io/otel"
)

// ErrorResponse represents a response containing an error message.
type ErrorResponse struct {
	Error       string            `json:"error"`
	Validations validation.Errors `json:"validations,omitempty"`
}

func ErrResponse(ctx context.Context, w http.ResponseWriter, msg string, err error) {
	resp := ErrorResponse{Error: msg}
	status := http.StatusInternalServerError

	var ierr *xerrors.Error
	if !errors.As(err, &ierr) {
		resp.Error = "internal error"
	} else {
		switch ierr.Code() {
		case xerrors.ErrorCodeNotFound:
			status = http.StatusNotFound
		case xerrors.ErrorCodeInvalidArgument:
			status = http.StatusBadRequest

			var verrors validation.Errors
			if errors.As(ierr, &verrors) {
				resp.Validations = verrors
			}
		case xerrors.ErrorCodeUnknown:
			fallthrough
		default:
			status = http.StatusInternalServerError
		}
	}

	if err != nil {
		_, span := otel.Tracer("").Start(ctx, "renderErrorResponse")
		defer span.End()
		span.RecordError(err)
	}

	Response(w, resp, status)
}

func Response(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(status)

	_, _ = w.Write(content)
}
