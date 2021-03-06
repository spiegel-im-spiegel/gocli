#!/bin/bash
go mod verify || exit 1
go mod tidy -v || exit 1
go list -json -m all | docker run --rm -i sonatypecommunity/nancy:latest sleuth -n || exit 1
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.37.1 golangci-lint run --enable gosec ./... || exit 1
go test ./...
