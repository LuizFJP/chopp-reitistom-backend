package application_test

import (
	"chopp-reitistom-backend/application"
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/infrastructure/persistence/model"
	"testing"

	"github.com/google/uuid"
)

type SpyAddressRepository struct {
	createWasCalled bool
	createdModel    *model.Address
	createdErr      error

	getByUUIDWasCalled bool
	getUUID            uuid.UUID
	getByUUIDModel     *model.Address
	getErr             error
}

func (ar *SpyAddressRepository) Create(address *model.Address) error {
	ar.createWasCalled = true
	ar.createdModel = address
	return ar.createdErr
}

func (ar *SpyAddressRepository) GetByUUID(uuid uuid.UUID) (*model.Address, error) {
	ar.getByUUIDWasCalled = true
	ar.getUUID = uuid
	return ar.getByUUIDModel, ar.getErr
}

func (ar *SpyAddressRepository) Update(address *model.Address) error { return nil }

func (ar *SpyAddressRepository) Delete(address *model.Address) error { return nil }

type SpyUserRepository struct {
	saveWasCalled   bool
	saveAffectedRow int64
	saveError       error

	getByEmailWasCalled bool
	getByEmailModel     *model.User
	getByEmailError     error

	getByUUIDWasCalled bool
	getByUUIDModel     *model.User
	getByUUIDError     error

	updateWasCalled   bool
	updateAffectedRow int64
	updateError       error

	deleteWasCalled   bool
	deleteAffectedRow int64
	deleteError       error
}

func (ur *SpyUserRepository) Save(user *model.User) (int64, error) {
	ur.saveWasCalled = true
	return ur.saveAffectedRow, ur.saveError
}

func (ur *SpyUserRepository) GetByEmail(email string) (*model.User, error) {
	ur.getByEmailWasCalled = true
	return ur.getByEmailModel, ur.getByEmailError
}

func (ur *SpyUserRepository) GetByUUID(uuid uuid.UUID) (*model.User, error) {
	ur.getByUUIDWasCalled = true
	return ur.getByUUIDModel, ur.getByUUIDError
}

func (ur *SpyUserRepository) Update(user *model.User) (int64, error) {
	ur.updateWasCalled = true
	return ur.updateAffectedRow, ur.updateError
}

func (ur *SpyUserRepository) Delete(user *model.User) (int64, error) {
	ur.deleteWasCalled = true
	return ur.deleteAffectedRow, ur.deleteError
}

func TestAddressUseCase_Create_Success(t *testing.T) {
	userRepositoryMock := SpyUserRepository{
		getByUUIDModel: &model.User{},
		getByUUIDError: nil,
	}

	addressRepositoryMock := SpyAddressRepository{
		createdModel: &model.Address{},
		createdErr:   nil,

		getByUUIDModel: &model.Address{},
		getErr:         nil,
	}

	addressUseCase := application.NewAddressUseCase(
		&addressRepositoryMock, &userRepositoryMock)

	got, err := addressUseCase.Create(&entity.Address{})

	if got == nil {
		t.Errorf("Could not save address. Got: %v", got)
	}

	if err != nil {
		t.Errorf("wanted call")
	}

	if !userRepositoryMock.getByUUIDWasCalled {
		t.Errorf("getByUUID was supposed to be called for userRepository. got %v", userRepositoryMock.getByUUIDWasCalled)
	}

	if !addressRepositoryMock.createWasCalled {
		t.Errorf("getByUUID was supposed to be called for addressRepository. got %v", userRepositoryMock.getByUUIDWasCalled)
	}

	if !addressRepositoryMock.getByUUIDWasCalled {
		t.Errorf("getByUUID was supposed to be called for addressRepository. got %v", userRepositoryMock.getByUUIDWasCalled)
	}
}
