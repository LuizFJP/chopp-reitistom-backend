package application

import (
	"chopp-reitistom-backend/config"
	"chopp-reitistom-backend/domain/entity"
	domainModel "chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/domain/repository"
	"chopp-reitistom-backend/infrastructure/mapper"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUserCaseInterface interface {
	SignUp(user domainModel.User) error
	SignIn(user entity.SignIn) (*entity.Token, error)
}

type AuthUseCase struct {
	authConfig     *config.Auth
	userRepository repository.UserRepositoryInterface
}

func NewAuthUseCase(
	authConfig *config.Auth,
	userRepository repository.UserRepositoryInterface) *AuthUseCase {
	return &AuthUseCase{
		authConfig:     authConfig,
		userRepository: userRepository,
	}
}

func (auc *AuthUseCase) SignUp(user domainModel.User) error {
	if auc.CheckEmailAlreadyExists(user.Email) {
		return domainModel.Wrap(
			fmt.Errorf("%s: %w", errors.New("email already exists"),
				domainModel.ErrInvalid))
	}
	user.UUID = uuid.New()
	auc.hashPassword(&user)
	userModel := mapper.FromEntityToModel(&user)
	auc.userRepository.Save(userModel)

	return nil
}

func (auc *AuthUseCase) SignIn(userEntity entity.SignIn) (*entity.Token, error) {
	user, err := auc.userRepository.GetByEmail(userEntity.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("there's no one with this email")
	}

	userDomainModel := mapper.FromModelToEntity(user)

	if auc.isNotValidPassword(userDomainModel, userEntity.Password) {
		return nil, errors.New("email or password is incorrect")
	}

	token, err := CreateToken(auc.authConfig)

	if err != nil {
		return nil, err
	}

	return &entity.Token{Token: token}, nil

}

func (auc *AuthUseCase) CheckEmailAlreadyExists(email string) bool {
	result, err := auc.userRepository.GetByEmail(email)
	if err != nil {
		fmt.Println("Error getting user by email", err)
	}

	return result != nil
}

func CreateToken(authConfig *config.Auth) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,
		jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(authConfig.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(authConfig *config.Auth, tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return authConfig.SecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func (auc *AuthUseCase) isNotValidPassword(user *domainModel.User, password string) bool {
	incomingUser := &domainModel.User{Password: password}
	auc.hashPassword(incomingUser)
	return user.Password != incomingUser.Password
}

func (auc *AuthUseCase) hashPassword(user *domainModel.User) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error while hashing password", err)
	}
	user.Password = string(hashedPassword)
}
