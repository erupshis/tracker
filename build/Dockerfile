ARG GO_VERSION=1.21

FROM golang:${GO_VERSION} AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -a -o tracker ./cmd/main.go

FROM alpine:3.21.3

WORKDIR /app

COPY --from=builder /app/tracker tracker

EXPOSE 8080

CMD ["/app/tracker"]