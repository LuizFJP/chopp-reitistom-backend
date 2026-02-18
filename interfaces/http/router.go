package http

import (
	"chopp-reitistom-backend/interfaces/http/handlers"
	"net/http"

	"github.com/vardius/gorouter"
)

func NewRouter(
	userHandler *handlers.UserHandler,
	authHandler *handlers.AuthHandler,
	addressHandler *handlers.AddressHandler,
	suggestionHandler *handlers.SuggestionHandler,
	categoryHandler *handlers.CategoryHandler,
	ratingHandler *handlers.RatingHandler,
	ingredientHandler *handlers.IngredientHandler,
	orderHandler *handlers.OrderHandler,
	productHandler *handlers.ProductHandler,
) http.Handler {
	publicRouter := gorouter.New()
	privateRouter := gorouter.New(authHandler.BearerAuthHandler)

	publicRouter.POST("/signup", authHandler.SignUpUserHandler())
	publicRouter.POST("/signin", authHandler.SignInUserHandler())
	privateRouter.GET("/user/{id}", userHandler.GetUserHandler())
	privateRouter.PATCH("/user", userHandler.UpdateUserHandler())
	privateRouter.DELETE("/user/{id}", userHandler.DeleteUserHandler())

	publicRouter.POST("/address", addressHandler.CreateHandler())
	privateRouter.GET("/address/user/{userId}", addressHandler.GetHandler())
	privateRouter.PATCH("/address", addressHandler.UpdateHandler())
	privateRouter.DELETE("/address/{addressId}", addressHandler.DeleteHandler())

	publicRouter.POST("/category", categoryHandler.CreateHandler())
	privateRouter.GET("/category/{categoryId}", categoryHandler.GetUUIDHandler())
	privateRouter.PATCH("/category", categoryHandler.UpdateHandler())
	privateRouter.DELETE("/category/{categoryId}", categoryHandler.DeleteHandler())

	publicRouter.POST("/suggestion", suggestionHandler.CreateHandler())
	privateRouter.GET("/suggestion/all", suggestionHandler.GetAllHandler())
	privateRouter.GET("/suggestion/user/{userId}", suggestionHandler.GetAllByUserHandler())
	privateRouter.PATCH("/suggestion", suggestionHandler.UpdateHandler())
	privateRouter.DELETE("/suggestion/{suggestionId}", suggestionHandler.DeleteHandler())

	publicRouter.POST("/rating", ratingHandler.CreateHandler())
	privateRouter.GET("/rating/all", ratingHandler.GetAll())
	privateRouter.GET("/rating/product/{productId}", ratingHandler.GetAllByProductIdHandler())
	privateRouter.DELETE("/rating/{ratingId}", ratingHandler.Delete())

	publicRouter.POST("/ingredient", ingredientHandler.CreateHandler())
	privateRouter.GET("/ingredient/all", ingredientHandler.GetAll())
	privateRouter.PATCH("/ingredient", ingredientHandler.UpdateHandler())
	privateRouter.GET("/ingredient/product/{productId}", ingredientHandler.GetAllByProductIdHandler())
	privateRouter.DELETE("/ingredient/{ingredientId}", ingredientHandler.DeleteHandler())

	publicRouter.POST("/order", orderHandler.CreateHandler())
	privateRouter.GET("/order/all/user/{userId}", orderHandler.GetAllByUserUUIDHandler())
	privateRouter.PATCH("/order", orderHandler.UpdateHandler())
	privateRouter.GET("/order/{orderId}", orderHandler.GetByUUIDHandler())
	privateRouter.DELETE("/order/{ingredientId}", orderHandler.DeleteHandler())

	publicRouter.POST("/product", productHandler.CreateHandler())
	privateRouter.GET("/product/all", productHandler.GetAllHandler())
	privateRouter.PATCH("/product", productHandler.UpdateHandler())
	privateRouter.GET("/product/{productId}", productHandler.GetByUUIDHandler())
	privateRouter.DELETE("/product/{ingredientId}", productHandler.DeleteHandler())
	privateRouter.PATCH("/product/addIngredients", productHandler.AddIngredientsHandler())
	privateRouter.PATCH("/product/removeIngredients", productHandler.RemoveIngredientsHandler())
	privateRouter.GET("/product/{productId}", productHandler.GetQuantityHandler())

	mainRouter := gorouter.New()
	mainRouter.GET("/health", handlers.BuildLivenessHandler())

	mainRouter.Mount("/v1/auth", publicRouter)
	mainRouter.Mount("/v1", privateRouter)

	return mainRouter
}
