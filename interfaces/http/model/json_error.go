package model

import (
	"context"
	"net/http"
)

// JSONError returns data as json response
// uses JSON internally
func JSONError(ctx context.Context, w http.ResponseWriter, err error) error {
	if err == nil {
		panic("JSONError called for nil error")
	}

	httpError := NewHttpError(ctx, err)

	if err := JSON(w, httpError.Code, httpError); err != nil {
		return err
	}

	return nil
}

// MustJSONError returns data as json response
// will panic if unable to marshal error into JSON object
// uses JSONError internally
func MustJSONError(ctx context.Context, w http.ResponseWriter, err error) {
	if err := JSONError(ctx, w, err); err != nil {
		panic(err)
	}
}
