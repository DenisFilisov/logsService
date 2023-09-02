package repository

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/mongo"
)

type Logs interface {
	SaveLogs(message *kafka.Message)
}

type Repository struct {
	Logs
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		Logs: NewLogsMongo(db),
	}
}
