package kafka

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

type KafkaConsumer struct {
	MsgChannel chan *ckafka.Message
}

func NewKafkaConsumer(msgChannel chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChannel: msgChannel,
	}
}

func(k *KafkaConsumer) Consume(){
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaCOnsumerGroupId"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("ERROR consuming kafka message:" + err.Error())
	}
	topics := []string{os.Getenv(KafkaReadTopic)}
	c.SubscribeTopics(topics, nil)
	fmt.Println("-- Kafka consumer has been started --")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChannel <- msg
		}
	}
}
