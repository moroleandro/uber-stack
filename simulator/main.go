package main

import (
	"fmt"
	"log"
	"github.com/moroleandro/uber-stack/simulator/infra/kafka"
	"github.com/joho/godotenv"
	kafka2 "github.com/moroleandro/uber-stack/application/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR loading .env file in project :(")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}
