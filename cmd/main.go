package main

import (
	"log"
	"todo-app"
	"todo-app/pkg/hendler"
)

func main() {
	hendlers := new(hendler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", hendlers.InitRoutes()); err != nil {
		log.Fatal("error occurent while running http server: %s", err.Error())
	}
}
