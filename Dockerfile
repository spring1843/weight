FROM golang:1.14.7-buster AS builder

WORKDIR /go/src/spring1843/weight/

COPY . .

RUN CGO_ENABLED=0  go build -o /go/bin/weight

ENTRYPOINT ["/go/bin/weight"]
