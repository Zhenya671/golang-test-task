FROM golang:latest as builder

RUN mkdir -p $GOPATH/src
WORKDIR $GOPATH/src
ADD . .
ENV GO111MODULE=on

RUN go build -o /bin/app main.go

FROM ubuntu:latest

COPY --from=builder /bin/app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY migration/ /etc/migration/
COPY config/ /etc/configs/

EXPOSE 8000

CMD ["./app"]
