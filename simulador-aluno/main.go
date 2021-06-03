package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/thgvieira/code-delivery/infra/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka2 "github.com/thgvieira/code-delivery/application/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for message := range msgChan {
		go kafka2.Produce(message)
	}
}
