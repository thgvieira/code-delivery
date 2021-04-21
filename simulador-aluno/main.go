package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/thgvieira/code-delivery/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola terraquios...", "readtest", producer)

	log.Println("Init the main")

	for {
		_ = 1

		time.Sleep(1 * time.Second)
		kafka.Publish("Mensagem ...", "readtest", producer)
	}
}
