#!/bin/bash

echo "Iniciando os testes..."
docker exec -it goapp go mod tidy
docker exec -it goapp go test -v ./...
