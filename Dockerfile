FROM golang:1.22-bullseye AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/bin/order ./cmd/order
RUN CGO_ENABLED=0 go build -o /app/bin/payment ./cmd/payment
RUN CGO_ENABLED=0 go build -o /app/bin/delivery ./cmd/delivery
RUN CGO_ENABLED=0 go build -o /app/bin/gateway ./cmd/gateway

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/bin/ /usr/local/bin/
EXPOSE 50051 50052 50053 50054 8080 8081 8082 8083
ENTRYPOINT ["/usr/local/bin/gateway"]
