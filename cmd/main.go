package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Config file error %s,", err.Error())
	}

	db, err := repository.NewPostgresInstance(repository.DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.dbname"),
		SSLMode:  viper.GetString("database.ssl_mode"),
	})

	if err != nil {
		logrus.Fatalf("Database init error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Server running error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("cfg")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
