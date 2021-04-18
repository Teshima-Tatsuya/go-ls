FROM golang:1.16.3-alpine

WORKDIR /go/src/app

ADD ../cmd /go/src/app

RUN ["go run /go/src/app/main.go"]
