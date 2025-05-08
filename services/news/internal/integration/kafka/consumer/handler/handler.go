package handler

import "github.com/IBM/sarama"

type MessageHandler interface {
	HandleMessage(msg *sarama.ConsumerMessage) error
}
