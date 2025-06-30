FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/stress-test

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

ENTRYPOINT ["/app/main"]
