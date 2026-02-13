package main

import (
	"chopp-reitistom-backend/application"
	"chopp-reitistom-backend/config"
	"chopp-reitistom-backend/infrastructure/persistence/postgres"
	httpInterface "chopp-reitistom-backend/interfaces/http"
	"chopp-reitistom-backend/interfaces/http/handlers"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configs := config.FromEnv()

	db, err := postgres.NewDB(configs.DB)
	defer db.Close()
	db.Automigrate()

	if err != nil {
		panic(err)
	}

	userApplication := application.NewUserUseCase(db.User)
	authApplication := application.NewAuthUseCase(&configs.Auth, db.User)
	addressApplication := application.NewAddressUseCase(db.Address, db.User)
	categoryApplication := application.NewCategoryUseCase(db.Category)
	suggestionApplication := application.NewSuggestionUseCase(db.Suggestion, db.User)
	ratingApplication := application.NewRatingUseCase(db.Rating, db.Product)
	ingredientApplication := application.NewIngredientUseCase(db.Ingredient, db.Product)

	userHandler := handlers.NewUserHandler(userApplication)
	authHandler := handlers.NewAuthHandler(&configs.Auth, authApplication)
	addressHandler := handlers.NewAddressHandler(addressApplication)
	suggestionHandler := handlers.NewSuggestionHandler(suggestionApplication)
	categoryHandler := handlers.NewCategoryHandler(categoryApplication)
	ratingHandler := handlers.NewRatingHandler(ratingApplication)
	ingredientHandler := handlers.NewIngredientHandler(ingredientApplication)

	router := httpInterface.NewRouter(
		userHandler,
		authHandler,
		addressHandler,
		suggestionHandler,
		categoryHandler,
		ratingHandler,
		ingredientHandler)

	go func() {
		if err = http.ListenAndServe(":8000", router); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Server is running on 8000")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	<-done
}
