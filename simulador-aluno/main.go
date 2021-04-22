package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/thgvieira/code-delivery/infra/kafka"

	kafka2 "github.com/thgvieira/code-delivery/simulador-aluno/application/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
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

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}
