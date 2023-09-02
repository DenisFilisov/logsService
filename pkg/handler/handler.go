package handler

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
	"logsService/pkg/service"
	"os"
)

type Handler struct {
	services *service.Service
	consumer *kafka.Consumer
}

func NewHandler(services *service.Service, consumer *kafka.Consumer) *Handler {
	return &Handler{services: services, consumer: consumer}
}

func (h *Handler) HandleMessage(consumer *kafka.Consumer) {
	run := true
	for run == true {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			h.services.Logs.SaveLogs(e)
		case kafka.Error:
			logrus.Error(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
		}
	}
}
