FROM golang:1.23-alpine

RUN apk update && apk add --no-cache git curl

WORKDIR /app

COPY . .

RUN go mod download
RUN go install github.com/air-verse/air@latest

CMD ["/go/bin/air"]
