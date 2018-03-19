FROM golang:1.9.2-alpine3.6 AS builder

RUN mkdir -p /go/src \
&& mkdir -p /go/bin \
&& mkdir -p /go/pkg

ENV GOPATH=/go

ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/src/restClient \
&& mkdir -p $GOPATH/src/github.com/YAWAL/ConfRESTcli/entities

ADD ./restClient $GOPATH/src/restClient
ADD entities $GOPATH/src/github.com/YAWAL/ConfRESTcli/entities

ADD ./vendor $GOPATH/src/vendor
ADD ./Gopkg.lock $GOPATH/src/
ADD ./Gopkg.toml $GOPATH/src/

WORKDIR $GOPATH/src/restClient

RUN go build -o $GOPATH/bin/restClient .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app

WORKDIR /app

COPY --from=builder /go/bin/restClient .

CMD ["./restClient"]

EXPOSE $CLIENT_PORT