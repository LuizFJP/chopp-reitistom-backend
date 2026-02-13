package handlers

import (
	"chopp-reitistom-backend/application"
	domainModel "chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/interfaces/http/model"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type RatingHandler struct {
	ratingUseCase application.RatingUseCaseInterface
}

func NewRatingHandler(
	ratingUseCase application.RatingUseCaseInterface,
) *RatingHandler {
	return &RatingHandler{ratingUseCase}
}

func (rh *RatingHandler) CreateHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var rating domainModel.Rating
		json.NewDecoder(r.Body).Decode(&rating)

		ratingResponse, err := rh.ratingUseCase.Create(&rating)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, ratingResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (rh *RatingHandler) GetAllByProductIdHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		productUUID := r.FormValue("productId")

		if productUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}

		ratingsResponse, err := rh.ratingUseCase.GetAllByProductId(uuid.MustParse(productUUID))

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusOK, ratingsResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (rh *RatingHandler) GetAll() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		ratingsResponse, err := rh.ratingUseCase.GetAll()

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusOK, ratingsResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (rh *RatingHandler) Delete() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		ratingUUID := r.FormValue("ratingId")
		if ratingUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}
		uuidConverted := rh.ratingUseCase.Delete(uuid.MustParse(ratingUUID))

		if err := model.JSON(w, http.StatusNoContent, uuidConverted); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}
