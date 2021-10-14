FROM golang:1.17.2-alpine3.14

RUN mkdir /app
ADD app /app

WORKDIR /app
RUN go build -o main .
