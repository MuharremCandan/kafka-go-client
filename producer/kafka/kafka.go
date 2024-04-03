package mb

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaCfg struct {
	Server   string
	ClientID string
	Acks     string
}

type kafkaCli struct {
	kafkaProducer *kafka.Producer
}

func KafkCli(cfg KafkaCfg) (*kafkaCli, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Server,
		"client.id":         cfg.ClientID,
		"acks":              cfg.Acks,
	})
	if err != nil {
		return nil, err
	}
	return &kafkaCli{
		kafkaProducer: p,
	}, nil
}

func (kc *kafkaCli) SendMessage(message []byte, topic string) (string, error) {

	deliverych := make(chan kafka.Event, 10000)
	err := kc.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	},
		deliverych,
	)
	if err != nil {
		return "", err
	}

	e := <-deliverych
	m := e.(*kafka.Message)
	close(deliverych)
	if m.TopicPartition.Error != nil {
		return "", m.TopicPartition.Error
	} else {
		return fmt.Sprintf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset), nil
	}

}
