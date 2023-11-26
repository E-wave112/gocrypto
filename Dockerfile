# syntax=docker/dockerfile:1

FROM golang:1.18

ENV REDIS_URL = ${REDIS_URL}

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /gocrypto-server

CMD [ "/gocrypto-server" ]