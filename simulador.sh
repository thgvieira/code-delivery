#!/bin/bash
cd simulador-aluno
docker-compose down
docker-compose up -d
winpty docker exec -it simulator go run main.go