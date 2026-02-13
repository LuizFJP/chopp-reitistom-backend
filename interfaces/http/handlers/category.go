package handlers

import (
	"chopp-reitistom-backend/application"
	domainModel "chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/interfaces/http/model"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type CategoryHandler struct {
	categoryUseCase application.CategoryUseCaseInterface
}

func NewCategoryHandler(categoryUseCase application.CategoryUseCaseInterface) *CategoryHandler {
	return &CategoryHandler{categoryUseCase}
}

func (ch *CategoryHandler) CreateHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var category domainModel.Category
		json.NewDecoder(r.Body).Decode(&category)

		categoryResponse, err := ch.categoryUseCase.Create(&category)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, categoryResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ch *CategoryHandler) GetUUIDHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		categoryUUID := r.FormValue("categoryId")

		if categoryUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}

		addressResponse, err := ch.categoryUseCase.Get(uuid.MustParse(categoryUUID))

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusOK, addressResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ch *CategoryHandler) UpdateHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var category domainModel.Category
		json.NewDecoder(r.Body).Decode(&category)

		categoryResponse, err := ch.categoryUseCase.Update(&category)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, categoryResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ch *CategoryHandler) DeleteHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		categoryUUID := r.FormValue("categoryId")
		if categoryUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}
		uuidConverted := ch.categoryUseCase.Delete(uuid.MustParse(categoryUUID))

		if err := model.JSON(w, http.StatusNoContent, uuidConverted); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}
