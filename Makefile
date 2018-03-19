include .env
export

all: dependencies build

.PHONY: build
build:
	echo "Build"
	go build -o ${GOPATH}/src/github.com/YAWAL/ConfRESTcli/bin/restClient ./restClient

.PHONY: run
run:
	echo "Running client"
	echo ${PDB_HOST}
	go build -o ${GOPATH}/src/github.com/YAWAL/ConfRESTcli/bin/restClient ./restClient
	./bin/restClient

.PHONY: dependencies
dependencies:
	echo "Installing dependencies"
	dep ensure

install dep:
	echo    "Installing dep"
	curl    https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

.PHONY: tests
tests:
	echo "Tests"
	go test ./restclient

docker-build:
	docker build -t configrestclient . && docker run -p ${CLIENT_PORT}:${CLIENT_PORT} configrestclient