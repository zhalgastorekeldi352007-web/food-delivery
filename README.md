# Food Delivery Microservices

A production-ready Go food delivery system built with clean architecture, gRPC, NATS event-driven communication, PostgreSQL, Redis, OpenTelemetry, Prometheus, Grafana, and SMTP email notifications.

## Services

- `order` — manages orders, publishes `OrderCreated`, consumes `PaymentProcessed`, `DeliveryAssigned`
- `payment` — processes payments, consumes `OrderCreated`, publishes `PaymentProcessed`
- `delivery` — assigns deliveries, consumes `PaymentProcessed`, publishes `DeliveryAssigned`
- `gateway` — API gateway with REST and gRPC auth, JWT, rate limiting, and request routing

## Key Features

- Clean architecture layers: entities, use cases, adapters, frameworks
- 12+ gRPC endpoints across services
- NATS event bus for asynchronous communication
- PostgreSQL with schemas and migrations
- Redis caching for active orders and sessions
- SMTP-based email notifications with HTML templates
- Prometheus metrics and Grafana observability
- Graceful shutdown and environment-based configuration

## Get Started

1. Create a `.env` file or export env vars for SMTP credentials.
2. Run Docker Compose:

```bash
docker compose up --build
```

3. Apply database migrations:

```bash
make migrate-up
```

4. Access services:

- Gateway REST: `http://localhost:8080`
- Prometheus: `http://localhost:9090`
- Grafana: `http://localhost:3000` (admin/admin)
- Jaeger: `http://localhost:16686`

## gRPC Endpoints

Available service actions include:

- Auth: `Register`, `Login`, `GetProfile`
- Order: `CreateOrder`, `UpdateOrderStatus`, `GetOrder`, `ListUserOrders`, `CancelOrder`
- Payment: `CreatePayment`, `GetPaymentStatus`
- Delivery: `AssignDeliveryPerson`, `UpdateDeliveryLocation`, `GetDeliveryStatus`, `ListDeliveries`

## Commands

```bash
make build
make test
make migrate-up
make migrate-down
```

## Monitoring

- Prometheus scrapes metrics from all services
- Grafana dashboards available in `http://localhost:3000`
- Jaeger collects spans from gRPC traffic

## Project Layout

- `cmd/order` — order service
- `cmd/payment` — payment service
- `cmd/delivery` — delivery service
- `cmd/gateway` — API gateway
- `internal/` — shared packages and business logic
- `proto/` — gRPC API definitions
- `migrations/` — golang-migrate SQL migrations
- `docker-compose.yml` — local environment with Postgres, Redis, NATS, Jaeger, Prometheus, Grafana

## Notes

- Replace SMTP credentials in environment variables before sending real email
- Use gateway REST endpoints for frontend integration
- The project uses a JSON gRPC codec to avoid external `protoc` dependencies for compilation
