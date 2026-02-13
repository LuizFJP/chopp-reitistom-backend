package application

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/domain/repository"
	"chopp-reitistom-backend/infrastructure/mapper"

	"github.com/google/uuid"
)

type RatingUseCaseInterface interface {
	Create(entity *entity.Rating) (*entity.Rating, error)
	Delete(uuid uuid.UUID) error
	GetAllByProductId(productUUID uuid.UUID) ([]*entity.Rating, error)
	GetAll() ([]*entity.Rating, error)
}

type RatingUseCase struct {
	ratingRepository  repository.RatingRepositoryInterface
	productRepository repository.ProductRepositoryInterface
}

func NewRatingUseCase(
	ratingRepository repository.RatingRepositoryInterface,
	productRepository repository.ProductRepositoryInterface,
) *RatingUseCase {
	return &RatingUseCase{
		ratingRepository,
		productRepository,
	}
}

func (ruc *RatingUseCase) Create(entity *entity.Rating) (*entity.Rating, error) {
	product, err := ruc.productRepository.GetByUUID(entity.ProductId)
	if err != nil {
		return nil, err
	}
	ratingModel := mapper.FromRatingEntityToModel(entity)
	ratingModel.UUID = uuid.New()
	ratingModel.ProductId = product.Id

	err = ruc.ratingRepository.Create(ratingModel)
	if err != nil {
		return nil, err
	}

	ratingSaved, err := ruc.ratingRepository.GetByUUID(ratingModel.UUID)
	if err != nil {
		return nil, err
	}

	return mapper.FromRatingModelToEntity(ratingSaved), nil
}

func (ruc *RatingUseCase) Delete(uuid uuid.UUID) error {
	rating, err := ruc.ratingRepository.GetByUUID(uuid)
	if err != nil {
		return err
	}

	if err = ruc.ratingRepository.Delete(rating); err != nil {
		return err
	}

	return nil
}

func (ruc *RatingUseCase) GetAllByProductId(productUUID uuid.UUID) ([]*entity.Rating, error) {
	product, err := ruc.productRepository.GetByUUID(productUUID)
	if err != nil {
		return nil, err
	}

	rating, err := ruc.ratingRepository.GetAllByProduct(product.Id)
	if err != nil {
		return nil, err
	}

	return mapper.FromRatingModelToEntityArray(rating), nil
}

func (ruc *RatingUseCase) GetAll() ([]*entity.Rating, error) {
	ratings, err := ruc.ratingRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return mapper.FromRatingModelToEntityArray(ratings), nil
}
