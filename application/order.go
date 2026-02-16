package application

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/domain/repository"
	"chopp-reitistom-backend/infrastructure/mapper"
	"fmt"

	"github.com/google/uuid"
)

type OrderUseCaseInterface interface {
	Create(entity *entity.Order) (*entity.Order, error)
	Update(entity *entity.Order) (*entity.Order, error)
	GetByUUID(orderUUID uuid.UUID) (*entity.Order, error)
	GetAllByUserUUID(userUUID uuid.UUID) ([]*entity.Order, error)
	Delete(orderUUID uuid.UUID) error
}

type OrderUseCase struct {
	orderRepository   repository.OrderRepositoryInterface
	productRepository repository.ProductRepositoryInterface
	userRepository    repository.UserRepositoryInterface
}

func NewOrderUseCase(
	orderRepository repository.OrderRepositoryInterface,
	productRepository repository.ProductRepositoryInterface,
	userRepository repository.UserRepositoryInterface,
) *OrderUseCase {
	return &OrderUseCase{
		orderRepository:   orderRepository,
		productRepository: productRepository,
		userRepository:    userRepository,
	}
}

func (ouc *OrderUseCase) Create(entity *entity.Order) (*entity.Order, error) {
	orderModel := mapper.FromOrderEntityToModel(entity)
	orderModel.UUID = uuid.New()
	productsModel, err := ouc.productRepository.GetManyByUUID(entity.ProductsUUID)
	if err != nil {
		return nil, err
	}

	orderModel.Products = productsModel
	userModel, err := ouc.userRepository.GetByUUID(entity.UserUUID)
	if err != nil {
		return nil, err
	}

	orderModel.UserId = userModel.Id
	if err = ouc.orderRepository.Create(orderModel); err != nil {
		return nil, err
	}

	orderCreated, err := ouc.orderRepository.GetByUUID(orderModel.UUID)
	if err != nil {
		return nil, err
	}

	entityCreated, err := mapper.FromOrderModelToEntity(orderCreated)
	if err != nil {
		return nil, err
	}

	return entityCreated, nil
}

func (ouc *OrderUseCase) Update(entity *entity.Order) (*entity.Order, error) {
	return nil, fmt.Errorf("not implemented yet")
}

func (ouc *OrderUseCase) GetByUUID(orderUUID uuid.UUID) (*entity.Order, error) {
	orderModel, err := ouc.orderRepository.GetByUUID(orderUUID)
	if err != nil {
		return nil, err
	}

	orderEntity, err := mapper.FromOrderModelToEntity(orderModel)
	if err != nil {
		return nil, err
	}

	return orderEntity, nil
}

func (ouc *OrderUseCase) GetAllByUserUUID(userUUID uuid.UUID) ([]*entity.Order, error) {
	userModel, err := ouc.userRepository.GetByUUID(userUUID)
	if err != nil {
		return nil, err
	}

	ordersModel, err := ouc.orderRepository.GetManyByUserId(userModel.Id)
	ordersEntity, err := mapper.FromOrderModelToEntityArray(ordersModel)
	if err != nil {
		return ordersEntity, err
	}

	return ordersEntity, nil
}

func (ouc *OrderUseCase) Delete(orderUUID uuid.UUID) error {
	orderModel, err := ouc.orderRepository.GetByUUID(orderUUID)
	if err != nil {
		return err
	}

	if err = ouc.orderRepository.Delete(orderModel); err != nil {
		return err
	}

	return nil
}
