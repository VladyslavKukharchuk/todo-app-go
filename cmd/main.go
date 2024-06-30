package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"todo-app-go"
	"todo-app-go/pkg/handler"
	"todo-app-go/pkg/repository"
	"todo-app-go/pkg/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurent while running http server: %s", err.Error())
	}
}
