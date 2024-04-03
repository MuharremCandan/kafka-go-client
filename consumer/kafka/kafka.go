package mb

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaCfg struct {
	Server  string
	GroupID string
}

type kafkaCli struct {
	kafkaConsumer *kafka.Consumer
}

func KafkCli(cfg KafkaCfg) (*kafkaCli, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Server,
		"group.id":          cfg.GroupID,
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		return nil, err
	}
	return &kafkaCli{
		kafkaConsumer: c,
	}, nil
}

func (kc *kafkaCli) ReadMessage(topic string) error {

	fmt.Printf("Start to consume message from : %s", topic)
	err := kc.kafkaConsumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	for {
		ev := kc.kafkaConsumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("Date:%s  Message on %s:\n%s\n",
				time.Now(), e.TopicPartition, string(e.Value))
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
		}
	}

}
