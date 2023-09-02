package service

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"logsService/pkg/repository"
)

type LogsService struct {
	repo repository.Logs
}

func NewLogsService(repo repository.Logs) *LogsService {
	return &LogsService{repo: repo}
}

func (l *LogsService) SaveLogs(message *kafka.Message) {
	l.repo.SaveLogs(message)
}
