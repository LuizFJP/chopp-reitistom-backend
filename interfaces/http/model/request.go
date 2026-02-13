package model

import (
	domainModel "chopp-reitistom-backend/domain/entity"
	"encoding/json"
	"net/http"
)

type ApplicationMethod[T any] func(*T) (*T, error)

func Request[T any](applicationMethod ApplicationMethod[T], statusCode int) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var body T
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return domainModel.Wrap(err)
		}

		response, err := applicationMethod(&body)
		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = JSON(w, statusCode, response); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return HandlerFunc(fn)
}
