package application

import (
	domainModel "chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/domain/repository"
	"chopp-reitistom-backend/infrastructure/mapper"

	"github.com/google/uuid"
)

type SuggestionUseCaseInterface interface {
	Create(entity *domainModel.Suggestion) (*domainModel.Suggestion, error)
	Delete(uuid uuid.UUID) error
	Update(entity *domainModel.Suggestion) (*domainModel.Suggestion, error)
	GetAllByUser(uuid uuid.UUID) ([]*domainModel.Suggestion, error)
	GetAll() ([]*domainModel.Suggestion, error)
}

type SuggestionUseCase struct {
	suggestionRepository repository.SuggestionRepositoryInterface
	userRepository       repository.UserRepositoryInterface
}

func NewSuggestionUseCase(
	suggestionRepository repository.SuggestionRepositoryInterface,
	userRepository repository.UserRepositoryInterface,
) *SuggestionUseCase {
	return &SuggestionUseCase{
		suggestionRepository,
		userRepository}
}

func (suc *SuggestionUseCase) Create(entity *domainModel.Suggestion) (*domainModel.Suggestion, error) {
	entity.UUID = uuid.New()
	model := mapper.FromSuggestionEntityToModel(entity)
	if err := suc.suggestionRepository.Create(model); err != nil {
		return nil, err
	}

	modelCreated, err := suc.suggestionRepository.GetByUUID(entity.UUID)
	if err != nil {
		return nil, err
	}

	return mapper.FromSuggestionModelToEntity(modelCreated), nil
}

func (suc *SuggestionUseCase) Delete(uuid uuid.UUID) error {
	suggestion, err := suc.suggestionRepository.GetByUUID(uuid)
	if err != nil {
		return err
	}

	if err = suc.suggestionRepository.Delete(suggestion); err != nil {
		return err
	}
	return nil
}

func (suc *SuggestionUseCase) Update(entity *domainModel.Suggestion) (*domainModel.Suggestion, error) {
	suggestionModel, err := suc.suggestionRepository.GetByUUID(entity.UUID)
	if err != nil {
		return nil, err
	}
	mapper.UpdateSuggestionFromEntityToModel(entity, suggestionModel)
	if err = suc.suggestionRepository.Update(suggestionModel); err != nil {
		return nil, err
	}

	suggestionUpdated, err := suc.suggestionRepository.GetByUUID(entity.UUID)

	return mapper.FromSuggestionModelToEntity(suggestionUpdated), nil
}

func (suc *SuggestionUseCase) GetAllByUser(userUUID uuid.UUID) ([]*domainModel.Suggestion, error) {
	user, err := suc.userRepository.GetByUUID(userUUID)
	if err != nil {
		return nil, err
	}

	suggestions, err := suc.suggestionRepository.GetAllByUser(user.Id)
	if err != nil {
		return nil, err
	}

	return mapper.FromSuggestionModelToEntityArray(suggestions), nil
}

func (suc *SuggestionUseCase) GetAll() ([]*domainModel.Suggestion, error) {
	suggestions, err := suc.suggestionRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return mapper.FromSuggestionModelToEntityArray(suggestions), nil
}
