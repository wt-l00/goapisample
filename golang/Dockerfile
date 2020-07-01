FROM golang:latest

RUN apt-get update -qq && \
    apt-get install -y default-mysql-client

WORKDIR /go/src/github.com/wt-l00/goapisample
COPY . .
RUN go build
