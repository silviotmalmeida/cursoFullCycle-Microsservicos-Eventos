#!/bin/bash

echo "Iniciando os testes..."
source runGoModTidy.sh
# docker exec -it goapp go test -v ./...
docker exec -it goapp go test ./...