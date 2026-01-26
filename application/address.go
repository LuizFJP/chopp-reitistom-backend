package application

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/domain/repository"
	"chopp-reitistom-backend/infrastructure/mapper"

	"github.com/google/uuid"
)

type AddressUseCaseInterface interface {
	Create(address *entity.Address) (*entity.Address, error)
	Delete(uuid uuid.UUID) error
	Get(uuid uuid.UUID) (*entity.Address, error)
	Update(address *entity.Address) (*entity.Address, error)
}

type AddressUseCase struct {
	addressRepository repository.AddressRepositoryInterface
	userRepository    repository.UserRepositoryInterface
}

func NewAddressUseCase(repository repository.AddressRepositoryInterface) *AddressUseCase {
	return &AddressUseCase{
		addressRepository: repository,
	}
}

func (auc *AddressUseCase) Create(entity *entity.Address) (*entity.Address, error) {
	user, err := auc.userRepository.GetByUUID(entity.UserUUID)

	if err != nil {
		return nil, err
	}

	entity.UUID = uuid.New()
	model := mapper.FromAddressEntityToModel(entity)
	model.UserId = user.Id
	if err := auc.addressRepository.Create(model); err != nil {
		return nil, err
	}

	modelCreated, err := auc.addressRepository.GetByUUID(model.UUID)
	if err != nil {
		return nil, err
	}

	return mapper.FromAddressModelToEntity(modelCreated), nil
}

func (auc *AddressUseCase) Delete(uuid uuid.UUID) error {
	address, err := auc.addressRepository.GetByUUID(uuid)
	if err != nil {
		return err
	}

	if err = auc.addressRepository.Delete(address); err != nil {
		return err
	}
	return nil
}

func (auc *AddressUseCase) Get(uuid uuid.UUID) (*entity.Address, error) {
	address, err := auc.addressRepository.GetByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return mapper.FromAddressModelToEntity(address), nil
}

func (auc *AddressUseCase) Update(address *entity.Address) (*entity.Address, error) {
	addressModel, err := auc.addressRepository.GetByUUID(address.UUID)
	if err != nil {
		return nil, err
	}

	mapper.UpdateAddressFromEntityToModel(address, addressModel)
	err = auc.addressRepository.Update(addressModel)

	if err != nil {
		return nil, err
	}

	addressUpdated, err := auc.addressRepository.GetByUUID(address.UUID)

	if err != nil {
		return nil, err
	}

	return mapper.FromAddressModelToEntity(addressUpdated), nil
}
