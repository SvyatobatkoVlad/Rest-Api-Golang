package main

import (
	server "github.com/SvyatobatkoVlad/Rest-Api-Golang"
	"log"
)

func main() {
	srv := new(server.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}

}
