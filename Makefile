BINARY_DIR := ./bin

.PHONY: build clean test migrate-up migrate-down compose

build:
	go build -o $(BINARY_DIR)/order ./cmd/order
	go build -o $(BINARY_DIR)/payment ./cmd/payment
	go build -o $(BINARY_DIR)/delivery ./cmd/delivery
	go build -o $(BINARY_DIR)/gateway ./cmd/gateway

clean:
	rm -rf $(BINARY_DIR)

test:
	go test ./... \
		-coverprofile=coverage.out

migrate-up:
	migrate -path ./migrations/order -database "${ORDER_DB_URL}" up
	migrate -path ./migrations/payment -database "${PAYMENT_DB_URL}" up
	migrate -path ./migrations/delivery -database "${DELIVERY_DB_URL}" up
	migrate -path ./migrations/gateway -database "${GATEWAY_DB_URL}" up

migrate-down:
	migrate -path ./migrations/order -database "${ORDER_DB_URL}" down
	migrate -path ./migrations/payment -database "${PAYMENT_DB_URL}" down
	migrate -path ./migrations/delivery -database "${DELIVERY_DB_URL}" down
	migrate -path ./migrations/gateway -database "${GATEWAY_DB_URL}" down

compose:
	docker compose up --build
