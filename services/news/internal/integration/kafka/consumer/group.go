package consumer

import (
	"github.com/IBM/sarama"
	"github.com/crypto-pulse/news/internal/integration/kafka/consumer/handler"
	"log"
)

type Consumer struct {
	consumerGroup sarama.ConsumerGroup
	topic         string
	handler       handler.MessageHandler
}

// TODO need to move all parameters to .yml
func NewConsumerGroup(brokers []string, topic, groupId string, handler handler.MessageHandler) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Return.Errors = true

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupId, config)
	if err != nil {
		return nil, err
	}

	go func() {
		for err := range consumerGroup.Errors() {
			log.Println(err)
		}
	}()

	return &Consumer{
		consumerGroup: consumerGroup,
		topic:         topic,
		handler:       handler,
	}, nil
}

func (c *Consumer) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if err := c.handler.HandleMessage(msg); err != nil {
			log.Println(err)
			continue
		}

		session.MarkMessage(msg, "")
	}

	return nil
}
