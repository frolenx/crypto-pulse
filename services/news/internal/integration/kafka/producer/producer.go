package producer

import (
	"github.com/IBM/sarama"
	"log"
	"time"
)

type Producer struct {
	asyncProducer sarama.AsyncProducer
	topic         string
}

// TODO need to move all parameters to .yml
func NewProducer(brokers []string, topic string, maxRetries int) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = maxRetries
	config.Producer.Retry.Backoff = 500 * time.Millisecond
	config.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	go func() {
		for err := range producer.Errors() {
			log.Println(err)
		}
	}()

	return &Producer{
		asyncProducer: producer,
		topic:         topic,
	}, nil
}

func (p *Producer) Publish(newsId string, msg []byte) {
	p.asyncProducer.Input() <- &sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder(newsId),
		Value: sarama.ByteEncoder(msg),
	}
}
