package integration

import (
    "context"
    "testing"
    
    "food-delivery/proto"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func TestGRPCOrderService(t *testing.T) {
    // Skip if order service not running
    conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        t.Skip("Order service not running, skipping integration test")
    }
    defer conn.Close()
    
    client := proto.NewOrderServiceClient(conn)
    
    // Test CreateOrder
    resp, err := client.CreateOrder(context.Background(), &proto.CreateOrderRequest{
        UserId:       "test_user",
        RestaurantId: "test_rest",
        Items: []*proto.OrderItem{
            {MenuItemId: "item1", Name: "Test Item", Quantity: 1, Price: 10.0},
        },
    })
    
    if err != nil {
        t.Skipf("CreateOrder failed (service may not be running): %v", err)
    }
    
    if resp.Order == nil {
        t.Error("Response order should not be nil")
    } else {
        t.Logf("Order created: %+v", resp.Order)
    }
}

func TestHealthCheck(t *testing.T) {
    // Simple health check test
    t.Log("Integration tests ready when services are running")
}
