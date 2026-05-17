//go:build integration
// +build integration

package integration

import (
    "testing"
)

func TestGRPCOrderService(t *testing.T) {
    t.Skip("Integration tests require running services")
}

func TestHealthCheck(t *testing.T) {
    t.Log("Integration tests ready when services are running")
}
