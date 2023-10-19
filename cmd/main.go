package main

import (
	"log"

	rest "github.com/Ubludor/restAPI_on_GO"
	"github.com/Ubludor/restAPI_on_GO/pkg/handler"
	"github.com/Ubludor/restAPI_on_GO/pkg/repository"
	"github.com/Ubludor/restAPI_on_GO/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(rest.Server)

	if err := srv.Run("8080",handlers.InitRouters()); err != nil {
		log.Fatal(err)
	}

}
