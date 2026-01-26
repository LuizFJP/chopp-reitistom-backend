package mapper

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/infrastructure/persistence/model"
)

func FromModelToEntity(userModel *model.User) *entity.User {
	return &entity.User{
		Email: userModel.Email,
		Name:  userModel.Name,
	}
}

func FromEntityToModel(userEntity *entity.User) *model.User {
	return &model.User{
		Email:    userEntity.Email,
		Password: userEntity.Password,
		UUID:     userEntity.UUID,
		Name:     userEntity.Name,
	}
}

func UpdateFromEntityToModel(userModel *model.User, userEntity *entity.User) *model.User {
	if userEntity.Name != "" {
		userModel.Name = userEntity.Name
	}

	if userEntity.Password != "" {
		userModel.Password = userEntity.Password
	}

	return userModel
}
