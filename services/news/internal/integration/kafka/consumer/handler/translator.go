package handler

import "github.com/IBM/sarama"

type TranslatorHandler struct{}

// TODO need to implement it. Now it is a stub.
func (h *TranslatorHandler) HandleMessage(msg *sarama.ConsumerMessage) error {
	return nil
}
