package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	route2 "github.com/thgvieira/code-delivery/application/route"
	"github.com/thgvieira/code-delivery/infra/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("kafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
