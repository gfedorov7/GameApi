package main

import (
	_ "GameApi/cmd/docs"
	"GameApi/internal/handler"
	"GameApi/internal/repository"
	"GameApi/internal/service"
	"GameApi/pkg/db"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	dsn := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database, err := db.NewDB(dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	userRepo := repository.NewUserRepo(database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Get("/users", userHandler.GetUsers)
	r.Get("/users/{id}", userHandler.GetUser)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Println("Starting server on :" + port)
	http.ListenAndServe(":"+port, r)
}
