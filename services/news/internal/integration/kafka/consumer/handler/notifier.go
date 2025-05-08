package handler

import "github.com/IBM/sarama"

type NotifierHandler struct{}

// TODO need to implement it. Now it is a stub.
func (h *NotifierHandler) HandleMessage(msg *sarama.ConsumerMessage) error {
	return nil
}
