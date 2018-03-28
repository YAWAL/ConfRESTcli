include .env
export

all: dependencies build

.PHONY: build
build:
	echo "Build"
	go build -o ${GOPATH}/src/github.com/YAWAL/ConfRESTcli/bin/restclient ./restclient

.PHONY: run
run:
	echo "Running client"
	go build -o ${GOPATH}/src/github.com/YAWAL/ConfRESTcli/bin/restclient ./restclient
	./bin/restclient

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
	CC=$(which musl-gcc) go build --ldflags '-w -linkmode external -extldflags "-static"' -o ${GOPATH}/src/github.com/YAWAL/ConfRESTcli/bin/restclient ./restclient && \
	docker build -t configrestclient . && \
	docker run --net=${DOCKER_NET_DRIVER} -p ${CLIENT_PORT}:${CLIENT_PORT} --env-file .env configrestclient

docker-build-with-compose:
	CC=$(which musl-gcc) go build --ldflags '-w -linkmode external -extldflags "-static"' -o ${GOPATH}/src/github.com/YAWAL/ConfRESTcli/bin/restclient ./restclient && \
	docker build -t configrestclient . && \
	docker run --network=${DOCKER_COMPOSE_NET} -p ${CLIENT_PORT}:${CLIENT_PORT} --env-file .env configrestclient