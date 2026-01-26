package application

import (
	domainModel "chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/domain/repository"
	"chopp-reitistom-backend/infrastructure/mapper"
	"fmt"

	"github.com/google/uuid"
)

type UserUseCaseInterface interface {
	Update(user domainModel.User) *domainModel.User
	Get(uuid uuid.UUID) *domainModel.User
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

func (uuc *UserUseCase) Update(user domainModel.User) *domainModel.User {
	userOldModel, err := uuc.userRepository.GetByUUID(user.UUID)

	if err != nil {
		fmt.Println("User not found", err)
	}

	userNewModel := mapper.UpdateFromEntityToModel(userOldModel, &user)
	_, err = uuc.userRepository.Update(userNewModel)

	if err != nil {
		fmt.Println("Error updating user", err)
	}

	userEntity := uuc.Get(userNewModel.UUID)

	return userEntity
}

func (uuc *UserUseCase) Get(uuid uuid.UUID) *domainModel.User {
	result, err := uuc.userRepository.GetByUUID(uuid)
	if err != nil {
		fmt.Println("Error getting user by UUID", err)
	}
	userEntity := mapper.FromModelToEntity(result)

	return userEntity
}

func (uuc *UserUseCase) Delete(uuid uuid.UUID) error {
	user := uuc.Get(uuid)
	userModel := mapper.FromEntityToModel(user)
	_, err := uuc.userRepository.Delete(userModel)

	if err != nil {
		fmt.Println("Error deleting user", err)
		return err
	}

	return nil
}
