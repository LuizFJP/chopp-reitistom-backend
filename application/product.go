package application

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/domain/repository"
	"chopp-reitistom-backend/infrastructure/mapper"
	"errors"

	"github.com/google/uuid"
)

type ProductUseCaseInterface interface {
	Create(product *entity.Product) (*entity.Product, error)
	Update(product *entity.Product) (*entity.Product, error)
	GetByUUID(uuid uuid.UUID) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	Delete(uuid uuid.UUID) error
	AddIngredients(productUUID uuid.UUID, ingredientUUID []uuid.UUID) error
	RemoveIngredients(productUUID uuid.UUID, ingredientUUID []uuid.UUID) error
	GetQuantity(productUUID uuid.UUID) (int, error)
}

type ProductUseCase struct {
	productRepository    repository.ProductRepositoryInterface
	itemRepository       repository.ItemRepositoryInterface
	ingredientRepository repository.IngredientRepositoryInterface
	categoryRepository   repository.CategoryRepositoryInterface
}

func NewProductUseCase(
	productRepository repository.ProductRepositoryInterface,
	itemRepository repository.ItemRepositoryInterface,
	ingredientRepository repository.IngredientRepositoryInterface,
	categoryRepository repository.CategoryRepositoryInterface,
) *ProductUseCase {
	return &ProductUseCase{
		productRepository:    productRepository,
		itemRepository:       itemRepository,
		ingredientRepository: ingredientRepository,
		categoryRepository:   categoryRepository,
	}
}

func (puc *ProductUseCase) Create(product *entity.Product) (*entity.Product, error) {
	ingredientUUIDs := mapper.ExtractIngredientUUIDs(product.Ingredients)
	ingredients, err := puc.ingredientRepository.GetManyByUUID(ingredientUUIDs)
	if err != nil {
		return nil, err
	}

	categoryModel, err := puc.categoryRepository.GetByUUID(product.CategoryUUID)
	if err != nil {
		return nil, err
	}

	productModel := mapper.FromProductEntityToModel(product)
	productModel.UUID = uuid.New()
	productModel.Ingredients = ingredients
	productModel.CategoryId = categoryModel.Id

	if err = puc.productRepository.Create(productModel); err != nil {
		return nil, err
	}

	productCreated, err := puc.productRepository.GetByUUID(productModel.UUID)

	if err != nil {
		return nil, err
	}

	items := mapper.CreateItemsModelFromQuantity(product.Quantity, productCreated.Id)
	quantityCreated, err := puc.itemRepository.CreateBatch(items, product.Quantity)

	if err != nil {
		return nil, err
	}

	if quantityCreated != product.Quantity {
		return nil, errors.New("A quantidade informada de produtos não foi criada corretamente")
	}

	productEntity := mapper.FromProductModelToEntity(productCreated)
	productEntity.Quantity = quantityCreated

	return productEntity, nil
}

func (puc *ProductUseCase) Update(product *entity.Product) (*entity.Product, error) {
	productModel, err := puc.productRepository.GetByUUID(product.UUID)
	if err != nil {
		return nil, err
	}

	mapper.UpdateProductFromEntityToModel(product, productModel)
	if err = puc.productRepository.Update(productModel); err != nil {
		return nil, err
	}

	productUpdated, err := puc.productRepository.GetByUUID(product.UUID)
	if err != nil {
		return nil, err
	}

	return mapper.FromProductModelToEntity(productUpdated), nil
}

func (puc *ProductUseCase) GetByUUID(uuid uuid.UUID) (*entity.Product, error) {
	product, err := puc.productRepository.GetByUUID(uuid)
	if err != nil {
		return nil, err
	}

	return mapper.FromProductModelToEntity(product), nil

}
func (puc *ProductUseCase) GetAll() ([]*entity.Product, error) {
	products, err := puc.productRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return mapper.FromProductFromModelToEntityArray(products), nil

}

func (puc *ProductUseCase) Delete(uuid uuid.UUID) error {
	product, err := puc.productRepository.GetByUUID(uuid)
	if err != nil {
		return err
	}

	if err = puc.productRepository.Delete(product); err != nil {
		return err
	}

	return nil

}

func (puc *ProductUseCase) AddIngredients(productUUID uuid.UUID, ingredientUUID []uuid.UUID) error {
	ingredients, err := puc.ingredientRepository.GetManyByUUID(ingredientUUID)
	if err != nil {
		return err
	}

	productModel, err := puc.productRepository.GetByUUID(productUUID)
	if err != nil {
		return err
	}

	if err = puc.productRepository.AddIngredients(productModel, ingredients); err != nil {
		return err
	}

	return nil
}
func (puc *ProductUseCase) RemoveIngredients(productUUID uuid.UUID, ingredientUUID []uuid.UUID) error {
	ingredients, err := puc.ingredientRepository.GetManyByUUID(ingredientUUID)
	if err != nil {
		return err
	}

	productModel, err := puc.productRepository.GetByUUID(productUUID)
	if err != nil {
		return err
	}

	if err = puc.productRepository.RemoveIngredients(productModel, ingredients); err != nil {
		return err
	}

	return nil
}
func (puc *ProductUseCase) GetQuantity(productUUID uuid.UUID) (int, error) {
	productModel, err := puc.productRepository.GetByUUID(productUUID)
	if err != nil {
		return 0, err
	}

	quantity, err := puc.itemRepository.GetQuantity(productModel.Id)
	if err != nil {
		return 0, err
	}

	return quantity, nil
}
