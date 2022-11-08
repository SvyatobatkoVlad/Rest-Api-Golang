package main

import (
	server "github.com/SvyatobatkoVlad/Rest-Api-Golang"
	handler "github.com/SvyatobatkoVlad/Rest-Api-Golang/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}
