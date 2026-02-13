package application

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/domain/repository"
	"chopp-reitistom-backend/infrastructure/mapper"

	"github.com/google/uuid"
)

type CategoryUseCaseInterface interface {
	Create(category *entity.Category) (*entity.Category, error)
	Delete(uuid uuid.UUID) error
	Get(uuid uuid.UUID) (*entity.Category, error)
	Update(category *entity.Category) (*entity.Category, error)
}

type CategoryUseCase struct {
	categoryRepository repository.CategoryRepositoryInterface
}

func NewCategoryUseCase(
	categoryRepository repository.CategoryRepositoryInterface,
) *CategoryUseCase {
	return &CategoryUseCase{
		categoryRepository,
	}
}

func (cuc *CategoryUseCase) Create(category *entity.Category) (*entity.Category, error) {
	category.UUID = uuid.New()
	model := mapper.FromCategoryEntityToModel(category)
	if err := cuc.categoryRepository.Create(model); err != nil {
		return nil, err
	}

	modelCreated, err := cuc.categoryRepository.GetByUUID(category.UUID)
	if err != nil {
		return nil, err
	}

	return mapper.FromCategoryModelToEntity(modelCreated), nil
}

func (cuc *CategoryUseCase) Delete(uuid uuid.UUID) error {
	category, err := cuc.categoryRepository.GetByUUID(uuid)
	if err != nil {
		return err
	}

	if err = cuc.categoryRepository.Delete(category); err != nil {
		return err
	}
	return nil
}

func (cuc *CategoryUseCase) Get(uuid uuid.UUID) (*entity.Category, error) {
	category, err := cuc.categoryRepository.GetByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return mapper.FromCategoryModelToEntity(category), nil
}

func (cuc *CategoryUseCase) Update(category *entity.Category) (*entity.Category, error) {
	categoryModel, err := cuc.categoryRepository.GetByUUID(category.UUID)
	if err != nil {
		return nil, err
	}

	mapper.UpdateCategoryFromEntityToModel(category, categoryModel)
	err = cuc.categoryRepository.Update(categoryModel)

	if err != nil {
		return nil, err
	}

	categoryUpdated, err := cuc.categoryRepository.GetByUUID(category.UUID)

	if err != nil {
		return nil, err
	}

	return mapper.FromCategoryModelToEntity(categoryUpdated), nil
}
