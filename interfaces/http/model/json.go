package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// ServeHTTP calls f(w, r) and handles error
func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		MustJSONError(r.Context(), w, err)
	}
}

func JSON(w http.ResponseWriter, statusCode int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	if payload == nil {
		_, err := w.Write([]byte("{}"))
		return err
	}

	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
	}

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(true)
	encoder.SetIndent("", "")

	if err := encoder.Encode(payload); err != nil {
		return err
	}

	Flush(w)

	return nil
}

func MustJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	if err := JSON(w, statusCode, payload); err != nil {
		panic(err)
	}
}

func NotFound() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		httpError := &HttpError{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Route %s %s", r.URL.Path, http.StatusText(http.StatusNotFound)),
		}

		return JSON(w, httpError.Code, httpError)
	}

	return HandlerFunc(fn)
}

func NotAllowed() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		httpError := &HttpError{
			Code:    http.StatusMethodNotAllowed,
			Message: http.StatusText(http.StatusMethodNotAllowed),
		}

		return JSON(w, httpError.Code, httpError)
	}

	return HandlerFunc(fn)
}
