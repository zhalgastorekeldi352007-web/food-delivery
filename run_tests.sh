#!/bin/bash

echo "================================"
echo "Running Test Suite"
echo "================================"

echo -e "\n1. Testing Order Service..."
go test ./internal/order/... -v -cover

echo -e "\n2. Testing Infrastructure..."
go test ./internal/infra/... -v -cover

echo -e "\n3. Testing Auth..."
go test ./internal/auth/... -v -cover

echo -e "\n4. Overall Coverage..."
go test ./internal/... -cover

echo -e "\n================================"
echo "✅ All tests completed!"
echo "================================"
