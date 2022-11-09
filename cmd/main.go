package main

import (
	server "github.com/SvyatobatkoVlad/Rest-Api-Golang"
	handler "github.com/SvyatobatkoVlad/Rest-Api-Golang/pkg/handler"
	"github.com/SvyatobatkoVlad/Rest-Api-Golang/pkg/repository"
	"github.com/SvyatobatkoVlad/Rest-Api-Golang/pkg/service"
	"log"
)

func main() {
	//Uncle Bob
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}
