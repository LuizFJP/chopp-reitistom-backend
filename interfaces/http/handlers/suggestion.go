package handlers

import (
	"chopp-reitistom-backend/application"
	domainModel "chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/interfaces/http/model"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type SuggestionHandler struct {
	suggestionUseCase application.SuggestionUseCaseInterface
}

func NewSuggestionHandler(
	suggestionUseCase application.SuggestionUseCaseInterface,
) *SuggestionHandler {
	return &SuggestionHandler{suggestionUseCase}
}

func (sh *SuggestionHandler) CreateHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var suggestion domainModel.Suggestion
		json.NewDecoder(r.Body).Decode(&suggestion)

		suggestionResponse, err := sh.suggestionUseCase.Create(&suggestion)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, suggestionResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (sh *SuggestionHandler) UpdateHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var suggestion domainModel.Suggestion
		json.NewDecoder(r.Body).Decode(&suggestion)

		suggestionResponse, err := sh.suggestionUseCase.Update(&suggestion)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, suggestionResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (sh *SuggestionHandler) GetAllByUserHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		userUUID := r.FormValue("userId")

		if userUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}

		suggestionsResponse, err := sh.suggestionUseCase.GetAllByUser(uuid.MustParse(userUUID))

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusOK, suggestionsResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (sh *SuggestionHandler) GetAllHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		suggestionsResponse, err := sh.suggestionUseCase.GetAll()

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusOK, suggestionsResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (sh *SuggestionHandler) DeleteHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		suggestionUUID := r.FormValue("suggestionId")
		if suggestionUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}
		uuidConverted := sh.suggestionUseCase.Delete(uuid.MustParse(suggestionUUID))

		if err := model.JSON(w, http.StatusNoContent, uuidConverted); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}
