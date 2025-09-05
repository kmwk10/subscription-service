FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o subscription-service ./cmd/app

FROM alpine:3.18

WORKDIR /root/

COPY --from=builder /app/subscription-service .

CMD ["./subscription-service"]
