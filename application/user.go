package application

import (
	domainModel "chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/domain/repository"
	"chopp-reitistom-backend/infrastructure/mapper"
	"fmt"

	"github.com/google/uuid"
)

type UserUseCaseInterface interface {
	Update(user domainModel.User) (*domainModel.User, error)
	Get(uuid uuid.UUID) (*domainModel.User, error)
	Delete(uuid uuid.UUID) error
}

type UserUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserUseCase(
	userRepository repository.UserRepositoryInterface,
) *UserUseCase {
	return &UserUseCase{userRepository}
}

func (uuc *UserUseCase) Update(user domainModel.User) (*domainModel.User, error) {
	userOldModel, err := uuc.userRepository.GetByUUID(user.UUID)

	if err != nil {
		fmt.Println("User not found", err)
	}

	userNewModel := mapper.UpdateFromEntityToModel(userOldModel, &user)
	_, err = uuc.userRepository.Update(userNewModel)

	if err != nil {
		fmt.Println("Error updating user", err)
		return nil, err
	}

	userEntity, err := uuc.Get(userNewModel.UUID)

	if err != nil {
		fmt.Println("User not found", err)
		return nil, err
	}

	return userEntity, nil
}

func (uuc *UserUseCase) Get(uuid uuid.UUID) (*domainModel.User, error) {
	result, err := uuc.userRepository.GetByUUID(uuid)
	if err != nil {
		fmt.Println("Error getting user by UUID", err)
		return nil, err
	}
	userEntity := mapper.FromModelToEntity(result)

	return userEntity, nil
}

func (uuc *UserUseCase) Delete(uuid uuid.UUID) error {
	user, err := uuc.Get(uuid)

	if err != nil {
		fmt.Println("Error getting user by UUID", err)
		return err
	}

	userModel := mapper.FromEntityToModel(user)
	_, err = uuc.userRepository.Delete(userModel)

	if err != nil {
		fmt.Println("Error deleting user", err)
		return err
	}

	return nil
}
