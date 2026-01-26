package handlers

import (
	"chopp-reitistom-backend/application"
	"chopp-reitistom-backend/config"
	domainModel "chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/interfaces/http/model"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type AuthHandler struct {
	authConfig  *config.Auth
	authUseCase application.AuthUserCaseInterface
}

func NewAuthHandler(authConfig *config.Auth,
	authUseCase application.AuthUserCaseInterface) *AuthHandler {
	return &AuthHandler{
		authConfig:  authConfig,
		authUseCase: authUseCase,
	}
}

func (au *AuthHandler) BearerAuthHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		rawToken := r.Header.Get("Authorization")
		tokenSplit := strings.Split(rawToken, "Bearer ")
		if len(tokenSplit) <= 1 {
			err := errors.New("missing authorization header")
			return model.JSON(w, http.StatusUnauthorized, err)
		}

		token := tokenSplit[1]

		if err := application.VerifyToken(token); err != nil {
			return model.JSON(w, http.StatusUnauthorized, err)

		}

		next.ServeHTTP(w, r)
		return nil
	}

	return model.HandlerFunc(fn)
}

func (ah *AuthHandler) SignUpUserHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var user domainModel.User
		json.NewDecoder(r.Body).Decode(&user)

		err := ah.authUseCase.SignUp(user)

		if err != nil {
			return domainModel.Wrap(err)
		}

		token, err := application.CreateToken(ah.authConfig)
		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, domainModel.Token{Token: token}); err != nil {
			return domainModel.Wrap(err)
		}

		return nil
	}

	return model.HandlerFunc(fn)
}

func (ah *AuthHandler) SignInUserHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return domainModel.Wrap(domainModel.ErrBodyIsMissing)
		}

		var user domainModel.SignIn
		json.NewDecoder(r.Body).Decode(&user)

		response, err := ah.authUseCase.SignIn(user)
		if err != nil {
			return domainModel.Wrap(err)
		}

		if err = model.JSON(w, http.StatusCreated, response); err != nil {
			return domainModel.Wrap(err)
		}

		return nil

	}

	return model.HandlerFunc(fn)
}
