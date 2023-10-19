package main

import (
	"log"

	rest "github.com/Ubludor/restAPI_on_GO"
	"github.com/Ubludor/restAPI_on_GO/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(rest.Server)

	if err := srv.Run("8080",handlers.InitRouters()); err != nil {
		log.Fatal(err)
	}

}
