#!/bin/bash
cd simulador-aluno
docker-compose down
docker-compose up -d
docker exec -it simulator go run main.go