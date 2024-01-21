FROM golang:1.21-alpine AS builder
LABEL authors="Arsmirnov"


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY backend .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app

FROM alpine:3.19

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD ["./main"]
