package handlers

import (
	"chopp-reitistom-backend/application"
	domainModel "chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/interfaces/http/model"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type UserHandler struct {
	userUseCase application.UserUseCaseInterface
}

func NewUserHandler(
	userUseCase application.UserUseCaseInterface,
) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (uh *UserHandler) GetUserHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		userUUID := r.FormValue("id")

		if userUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}

		result, err := uh.userUseCase.Get(uuid.MustParse(userUUID))

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err := model.JSON(w, http.StatusOK, result); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (uh *UserHandler) UpdateUserHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var user domainModel.User
		json.NewDecoder(r.Body).Decode(&user)

		result, err := uh.userUseCase.Update(domainModel.User{})

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err := model.JSON(w, http.StatusOK, result); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (uh *UserHandler) DeleteUserHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		userUUID := r.FormValue("id")
		if userUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}

		uuidConverted := uh.userUseCase.Delete(uuid.MustParse(userUUID))

		if err := model.JSON(w, http.StatusOK, uuidConverted); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}
