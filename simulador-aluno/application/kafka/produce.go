package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/thgvieira/code-delivery/application/route"
	"github.com/thgvieira/code-delivery/infra/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

//{"clientId":"1", "routeId":"1"}
//{"clientId":"2", "routeId":"2"}
//{"clientId":"3", "routeId":"3"}
func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()

	positions, error := route.ExportJsonPositions()

	if error != nil {
		log.Println(error.Error())
	}

	for _, position := range positions {
		fmt.Println("Position= " + position)
		kafka.Publish(position, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
