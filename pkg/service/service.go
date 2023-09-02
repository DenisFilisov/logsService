package service

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"logsService/pkg/repository"
)

type Logs interface {
	SaveLogs(message *kafka.Message)
}

type Service struct {
	Logs
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Logs: NewLogsService(repos),
	}
}
