package main

import (
	server "github.com/SvyatobatkoVlad/Rest-Api-Golang"
	handler "github.com/SvyatobatkoVlad/Rest-Api-Golang/pkg/handler"
	"github.com/SvyatobatkoVlad/Rest-Api-Golang/pkg/repository"
	"github.com/SvyatobatkoVlad/Rest-Api-Golang/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	if err := initConfig; err != nil {
		logrus.Fatalf("error reading config file: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	//Uncle Bob
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	//TODO viper viper.GetString("port")
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
