FROM golang:1.24.1-alpine AS builder

WORKDIR /build/news

COPY services/news/cmd ./cmd
COPY services/news/configs ./configs
COPY services/news/internal ./internal
COPY services/news/go.mod ./
COPY services/news/go.sum ./
COPY sdk ../../sdk

RUN go mod download

RUN go build -v -o news-service ./cmd/

FROM alpine:latest

WORKDIR /root

COPY --from=builder /build/news/news-service .
COPY --from=builder /build/news/configs ./configs

CMD ["./news-service"]