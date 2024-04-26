#!make

build:
	@ go build -o ./bin/app ./

run: build
	@ ./bin/app

test:
	@ go test -v ./... -count=1

coverage:
	@ go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out
