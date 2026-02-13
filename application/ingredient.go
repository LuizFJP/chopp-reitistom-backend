package application

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/domain/repository"
	"chopp-reitistom-backend/infrastructure/mapper"

	"github.com/google/uuid"
)

type IngredientUseCaseInterface interface {
	Create(entity *entity.Ingredient) (*entity.Ingredient, error)
	GetAllByProductId(productUUID uuid.UUID) ([]*entity.Ingredient, error)
	Update(entity *entity.Ingredient) (*entity.Ingredient, error)
	Delete(ingredientUUID uuid.UUID) error
	GetAll() ([]*entity.Ingredient, error)
}

type IngredientUseCase struct {
	ingredientRepository repository.IngredientRepositoryInterface
	productRepository    repository.ProductRepositoryInterface
}

func NewIngredientUseCase(
	ingredientRepository repository.IngredientRepositoryInterface,
	productRepository repository.ProductRepositoryInterface,
) *IngredientUseCase {
	return &IngredientUseCase{
		ingredientRepository,
		productRepository,
	}
}

func (iuc *IngredientUseCase) Create(entity *entity.Ingredient) (*entity.Ingredient, error) {
	product, err := iuc.productRepository.GetByUUID(entity.ProductId)
	if err != nil {
		return nil, err
	}
	ingredientModel := mapper.FromIngredientEntityToModel(entity)
	ingredientModel.UUID = uuid.New()
	ingredientModel.ProductId = product.Id

	err = iuc.ingredientRepository.Create(ingredientModel)
	if err != nil {
		return nil, err
	}

	ingredientSaved, err := iuc.ingredientRepository.GetByUUID(ingredientModel.UUID)
	if err != nil {
		return nil, err
	}

	return mapper.FromIngredientModelToEntity(ingredientSaved), nil
}

func (iuc *IngredientUseCase) GetAllByProductId(productUUID uuid.UUID) ([]*entity.Ingredient, error) {
	product, err := iuc.productRepository.GetByUUID(productUUID)
	if err != nil {
		return nil, err
	}

	ingredients, err := iuc.ingredientRepository.GetAllByProductId(product.Id)
	if err != nil {
		return nil, err
	}

	return mapper.FromIngredientModelToEntityArray(ingredients), nil
}

func (iuc *IngredientUseCase) Update(entity *entity.Ingredient) (*entity.Ingredient, error) {
	suggestionModel, err := iuc.ingredientRepository.GetByUUID(entity.UUID)
	if err != nil {
		return nil, err
	}
	mapper.UpdateIngredientFromEntityToModel(entity, suggestionModel)
	if err = iuc.ingredientRepository.Update(suggestionModel); err != nil {
		return nil, err
	}

	suggestionUpdated, err := iuc.ingredientRepository.GetByUUID(entity.UUID)

	return mapper.FromIngredientModelToEntity(suggestionUpdated), nil
}

func (iuc *IngredientUseCase) Delete(ingredientUUID uuid.UUID) error {
	rating, err := iuc.ingredientRepository.GetByUUID(ingredientUUID)
	if err != nil {
		return err
	}

	if err = iuc.ingredientRepository.Delete(rating); err != nil {
		return err
	}

	return nil
}

func (iuc *IngredientUseCase) GetAll() ([]*entity.Ingredient, error) {
	ingredients, err := iuc.ingredientRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return mapper.FromIngredientModelToEntityArray(ingredients), nil
}
