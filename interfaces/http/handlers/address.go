package handlers

import (
	"chopp-reitistom-backend/application"
	domainModel "chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/interfaces/http/model"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type AddressHandler struct {
	addressUseCase application.AddressUseCaseInterface
}

func NewAddressHandler(addressUseCase application.AddressUseCaseInterface) *AddressHandler {
	return &AddressHandler{addressUseCase}
}

func (ah *AddressHandler) CreateHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var address domainModel.Address
		json.NewDecoder(r.Body).Decode(&address)

		addressResponse, err := ah.addressUseCase.Create(&address)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, addressResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}
func (ah *AddressHandler) UpdateHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var address domainModel.Address
		json.NewDecoder(r.Body).Decode(&address)

		addressResponse, err := ah.addressUseCase.Update(&address)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, addressResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ah *AddressHandler) GetHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		userUUID := r.FormValue("userId")

		if userUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}

		addressResponse, err := ah.addressUseCase.Get(uuid.MustParse(userUUID))

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

func (ah *AddressHandler) DeleteHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		addressUUID := r.FormValue("addressId")
		if addressUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}
		uuidConverted := ah.addressUseCase.Delete(uuid.MustParse(addressUUID))

		if err := model.JSON(w, http.StatusNoContent, uuidConverted); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}
