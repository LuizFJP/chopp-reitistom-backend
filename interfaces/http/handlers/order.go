package handlers

import (
	"chopp-reitistom-backend/application"
	domainModel "chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/interfaces/http/model"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type OrderHandler struct {
	orderUseCase application.OrderUseCaseInterface
}

func NewOrderHandler(
	orderUseCase application.OrderUseCaseInterface,
) *OrderHandler {
	return &OrderHandler{orderUseCase}
}

func (oh *OrderHandler) CreateHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var order domainModel.Order
		json.NewDecoder(r.Body).Decode(&order)

		orderResponse, err := oh.orderUseCase.Create(&order)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, orderResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (oh *OrderHandler) UpdateHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var order domainModel.Order
		json.NewDecoder(r.Body).Decode(&order)

		orderResponse, err := oh.orderUseCase.Create(&order)

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, orderResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (oh *OrderHandler) GetByUUIDHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		orderUUID := r.FormValue("orderId")

		if orderUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}

		ordersResponse, err := oh.orderUseCase.GetByUUID(uuid.MustParse(orderUUID))

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusOK, ordersResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (oh *OrderHandler) GetAllByUserUUIDHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		userUUID := r.FormValue("userId")

		if userUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}

		ordersResponse, err := oh.orderUseCase.GetAllByUserUUID(uuid.MustParse(userUUID))

		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusOK, ordersResponse); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (oh *OrderHandler) DeleteHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		orderUUID := r.FormValue("orderId")
		if orderUUID == "" {
			return domainModel.Wrap(domainModel.ErrUUIDIsMissing)
		}
		err := oh.orderUseCase.Delete(uuid.MustParse(orderUUID))

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
