# syntax=docker/dockerfile:1

FROM golang:1.21-alpine3.18

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY pkg/ /app/pkg/
COPY server/ /app/server/
COPY cmd/ /app/cmd/

RUN CGO_ENABLED=0 GOOS=linux go build -o /gocrypto-server ./server

CMD ["/gocrypto-server"]