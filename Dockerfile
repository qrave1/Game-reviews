FROM golang:1.21-alpine AS builder
LABEL authors="Arsmirnov"


WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gameReviewBackend ./cmd/app

FROM alpine:3.19

COPY --from=builder /app/gameReviewBackend .

CMD ["./gameReviewBackend"]
