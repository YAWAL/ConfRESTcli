FROM golang:1.9.2-alpine3.6 AS build

RUN mkdir -p /go/src \
&& mkdir -p /go/bin \
&& mkdir -p /go/pkg

ENV GOPATH=/go

ENV PATH=$GOPATH/bin:$PATH

ENV CLIENT_PORT=8080

ENV SERVICE_HOST=localhost

ENV SERVICE_PORT=3000

RUN mkdir -p $GOPATH/src/restClient \
&& mkdir -p $GOPATH/src/github.com/YAWAL/ConfRESTcli/entitys \
&& mkdir -p $GOPATH/src/github.com/YAWAL/ConfRESTcli/api

ADD ./restClient $GOPATH/src/restClient
ADD ./entitys $GOPATH/src/github.com/YAWAL/ConfRESTcli/entitys
ADD ./api $GOPATH/src/github.com/YAWAL/ConfRESTcli/api

ADD ./vendor $GOPATH/src/vendor
ADD ./Gopkg.lock $GOPATH/src/
ADD ./Gopkg.toml $GOPATH/src/

WORKDIR $GOPATH/src/restClient

RUN go build -o main

CMD ["/go/src/restClient/main"]

EXPOSE $PORT