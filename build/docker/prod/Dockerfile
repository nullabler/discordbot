FROM golang:1.17.2-alpine3.14

RUN apk update && apk add git ffmpeg ca-certificates && update-ca-certificates

WORKDIR /app
COPY . /app

RUN go mod download && go build -o cmd/discordbot/main cmd/discordbot/main.go

ENTRYPOINT ["/app/cmd/discordbot/main"]
