package main

import (
	"log"
	"os"
	"time"

	"github.com/codeedu/imersaofsfc2-simulador/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	log.Println("topic: " + os.Getenv("KafkaProduceTopic"))
	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola terraquios...", "readtest", producer)

	log.Println("Init the main")

	for {
		_ = 1

		time.Sleep(1 * time.Second)
		kafka.Publish("Mensagem ...", "readtest", producer)
	}
}
