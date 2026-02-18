package handlers

import (
	"chopp-reitistom-backend/application"
	domainModel "chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/interfaces/http/model"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type ProductHandler struct {
	productUseCase application.ProductUseCaseInterface
}

func NewProductHandler(
	productUseCase application.ProductUseCaseInterface,
) *ProductHandler {
	return &ProductHandler{productUseCase}
}

func (ph *ProductHandler) CreateHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var product domainModel.Product
		json.NewDecoder(r.Body).Decode(&product)

		productResponse, err := ph.productUseCase.Create(&product)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, productResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ph *ProductHandler) UpdateHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var product domainModel.Product
		json.NewDecoder(r.Body).Decode(&product)

		productResponse, err := ph.productUseCase.Update(&product)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, productResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ph *ProductHandler) GetByUUIDHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		productUUID := r.FormValue("productId")

		if productUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}

		quantityResponse, err := ph.productUseCase.GetByUUID(uuid.MustParse(productUUID))

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusOK, quantityResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ph *ProductHandler) GetAllHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		quantityResponse, err := ph.productUseCase.GetAll()

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusOK, quantityResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ph *ProductHandler) DeleteHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		productUUID := r.FormValue("productId")
		if productUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}
		uuidConverted := ph.productUseCase.Delete(uuid.MustParse(productUUID))

		if err := model.JSON(w, http.StatusNoContent, uuidConverted); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ph *ProductHandler) AddIngredientsHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var productIngredients domainModel.ProductIngredients
		json.NewDecoder(r.Body).Decode(&productIngredients)

		err := ph.productUseCase.AddIngredients(productIngredients.ProductUUID, productIngredients.IngredientsUUID)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusNoContent, struct{}{}); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ph *ProductHandler) RemoveIngredientsHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var productIngredients domainModel.ProductIngredients
		json.NewDecoder(r.Body).Decode(&productIngredients)

		err := ph.productUseCase.RemoveIngredients(productIngredients.ProductUUID, productIngredients.IngredientsUUID)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusNoContent, struct{}{}); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ph *ProductHandler) GetQuantityHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		productUUID := r.FormValue("productId")

		if productUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}

		quantityResponse, err := ph.productUseCase.GetQuantity(uuid.MustParse(productUUID))

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusOK, quantityResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}
