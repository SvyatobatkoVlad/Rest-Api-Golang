package main

import (
	"context"
	server "github.com/SvyatobatkoVlad/Rest-Api-Golang"
	handler "github.com/SvyatobatkoVlad/Rest-Api-Golang/pkg/handler"
	"github.com/SvyatobatkoVlad/Rest-Api-Golang/pkg/repository"
	"github.com/SvyatobatkoVlad/Rest-Api-Golang/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//logrus set json format
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//if err := initConfig; err != nil {
	//	logrus.Fatalf("error reading config file: %s", err)
	//}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: os.Getenv("DB_PASSWORD"), //qwerty
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
	go func() {
		//TODO viper viper.GetString("port")
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error running http server: %s", err.Error())
		}
	}()

	logrus.Println("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error Shutdown failed: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error db.Close() failed: %s ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
