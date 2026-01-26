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
	authApplication := application.NewAuthUseCase(db.User)

	userHandler := handlers.NewUserHandler(userApplication)
	authHandler := handlers.NewAuthHandler(authApplication)

	router := httpInterface.NewRouter(userHandler, authHandler)

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
