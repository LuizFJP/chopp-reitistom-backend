package handlers

import (
	"chopp-reitistom-backend/application"
	domainModel "chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/interfaces/http/model"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type IngredientHandler struct {
	ingredientUseCase application.IngredientUseCaseInterface
}

func NewIngredientHandler(
	ingredientUseCase application.IngredientUseCaseInterface,
) *IngredientHandler {
	return &IngredientHandler{
		ingredientUseCase,
	}
}

func (ih *IngredientHandler) CreateHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var ingredient domainModel.Ingredient
		json.NewDecoder(r.Body).Decode(&ingredient)

		ingredientResponse, err := ih.ingredientUseCase.Create(&ingredient)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, ingredientResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ih *IngredientHandler) GetAllByProductIdHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		productUUID := r.FormValue("productId")

		if productUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}

		ingredientsResponse, err := ih.ingredientUseCase.GetAllByProductId(uuid.MustParse(productUUID))

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusOK, ingredientsResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

//func (ih *IngredientHandler) UpdateHandler() http.Handler {
//	fn := func(w http.ResponseWriter, r *http.Request) error {
//		if r.Body == nil {
//			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
//		}
//
//		var ingredient domainModel.Ingredient
//		json.NewDecoder(r.Body).Decode(&ingredient)
//
//		ingredientResponse, err := ih.ingredientUseCase.Update(&ingredient)
//
//		if err != nil {
//			return domainModel.Wrap(err)
//		}
//
//		if err = model.JSON(w, http.StatusCreated, ingredientResponse); err != nil {
//			return domainModel.Wrap(err)
//		}
//
//		return nil
//	}
//
//	return model.HandlerFunc(fn)
//}

func (ih *IngredientHandler) UpdateHandler() http.Handler {
	return model.Request[domainModel.Ingredient](ih.ingredientUseCase.Update, http.StatusCreated)
}

func (ih *IngredientHandler) DeleteHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		ingredientUUID := r.FormValue("ingredientId")
		if ingredientUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}
		uuidConverted := ih.ingredientUseCase.Delete(uuid.MustParse(ingredientUUID))

		if err := model.JSON(w, http.StatusNoContent, uuidConverted); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ih *IngredientHandler) GetAll() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		ingredientsResponse, err := ih.ingredientUseCase.GetAll()

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusOK, ingredientsResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}
