#!/bin/bash
cd apache-kafka
docker-compose down
docker-compose up -d
winpty docker exec -it apache-kafka_kafka_1 kafka-console-consumer --bootstrap-server=kafka:9092 --topic=route.new-position