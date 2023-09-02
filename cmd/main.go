package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"logsService/pkg/handler"
	"logsService/pkg/repository"
	"logsService/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error while initializing configuration %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:       os.Getenv("MONGO_DB_HOST"),
		Port:       os.Getenv("MONGO_DB_PORT"),
		Username:   os.Getenv("MONGO_DB_USERNAME"),
		Password:   os.Getenv("MONGO_DB_PASSWORD"),
		DataBase:   os.Getenv("MONGO_DB_DATABASE"),
		Collection: os.Getenv("MONGO_DB_COLLECTION"),
	})

	if err != nil {
		logrus.Fatalf("Can't initialize db: %s", err.Error())
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("kafka.bootstrap_servers"),
		"group.id":          viper.GetString("kafka.group-id"),
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	topics := make([]string, 0)
	topics = append(topics, viper.GetString("kafka.topics.consumer"))
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		logrus.Fatalf("Can't subscribe to topics %s", err)
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services, consumer)

	go func() {
		handlers.HandleMessage(consumer)
	}()
	logrus.Print("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App shutting down")

	if err := consumer.Close(); err != nil {
		logrus.Errorf("error while shutdown process: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
