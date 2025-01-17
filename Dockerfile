FROM golang:1.22-alpine3.20 as builder


WORKDIR  /app


COPY go.mod go.sum ./


RUN go mod download


COPY . .


RUN go build -o main .


CMD ["/app/main"]