package repository

import (
	"context"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogsMongo struct {
	db *mongo.Collection
}

func (l *LogsMongo) SaveLogs(message *kafka.Message) {

	var jsonData map[string]interface{}
	if err := json.Unmarshal(message.Value, &jsonData); err != nil {
		logrus.Error("Error parsing JSON message: %v\n", err)
	}

	log, err := l.db.InsertOne(context.Background(), jsonData)
	if err != nil {
		logrus.Info(err)
	}
	insertedID := log.InsertedID.(primitive.ObjectID).Hex()
	logrus.Infof("Log saved with ID: %s", insertedID)
}

func NewLogsMongo(db *mongo.Collection) *LogsMongo {
	return &LogsMongo{db: db}
}
