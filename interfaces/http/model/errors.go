package model

import (
	"chopp-reitistom-backend/domain/entity"
	"context"
	"errors"

	"net/http"
)

type HttpError struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id,omitempty"`
}

func NewHttpError(ctx context.Context, err error) *HttpError {
	code := http.StatusInternalServerError

	switch {
	case errors.Is(err, entity.ErrInvalid):
		code = http.StatusBadRequest
	case errors.Is(err, entity.ErrUnauthorized):
		code = http.StatusUnauthorized
	case errors.Is(err, entity.ErrForbidden):
		code = http.StatusForbidden
	case errors.Is(err, entity.ErrNotFound):
		code = http.StatusNotFound
	case errors.Is(err, entity.ErrTimeout):
		code = http.StatusRequestTimeout
	case errors.Is(err, entity.ErrTemporaryDisabled):
		code = http.StatusServiceUnavailable
	case errors.Is(err, entity.ErrBodyIsMissing):
		code = http.StatusBadRequest
	case errors.Is(err, entity.ErrUUIDIsMissing):
		code = http.StatusBadRequest
	case errors.Is(err, entity.ErrInternal):
		code = http.StatusInternalServerError
	}

	httpError := &HttpError{
		Code:    code,
		Message: err.Error(),
	}

	if m, ok := FromContext(ctx); ok {
		httpError.RequestID = m.TraceID
		m.Err = err
	}

	return httpError
}
