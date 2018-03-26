FROM alpine:latest

ADD ./bin  .

CMD ["./restclient"]

EXPOSE $CLIENT_PORT