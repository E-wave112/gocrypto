# syntax=docker/dockerfile:1

FROM golang:1.21-alpine3.18

ARG REDIS_URL

ENV REDIS_URL = ${REDIS_URL}

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY pkg/ /app/pkg/
COPY server/ /app/server/
COPY cmd/ /app/cmd/

EXPOSE 9000

# RUN pwd

RUN CGO_ENABLED=0 GOOS=linux go build -o /gocrypto-server ./server

CMD ["/gocrypto-server"]