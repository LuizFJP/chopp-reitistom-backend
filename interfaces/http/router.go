package http

import (
	"chopp-reitistom-backend/interfaces/http/handlers"
	"net/http"

	"github.com/vardius/gorouter"
)

func NewRouter(
	userHandler *handlers.UserHandler,
	authHandler *handlers.AuthHandler,
) http.Handler {
	publicRouter := gorouter.New()
	privateRouter := gorouter.New(authHandler.BearerAuthHandler)

	publicRouter.POST("/signup", authHandler.SignUpUserHandler())
	publicRouter.POST("/signin", authHandler.SignInUserHandler())
	privateRouter.GET("/user/{id}", userHandler.GetUserHandler())
	privateRouter.PATCH("/user", userHandler.UpdateUserHandler())
	privateRouter.DELETE("/user/{id}", userHandler.DeleteUserHandler())

	mainRouter := gorouter.New()
	mainRouter.GET("/health", handlers.BuildLivenessHandler())

	mainRouter.Mount("/v1/auth", publicRouter)
	mainRouter.Mount("/v1", privateRouter)

	return mainRouter
}
