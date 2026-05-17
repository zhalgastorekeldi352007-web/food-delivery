package order

import (
    "context"
    "testing"
    "go.uber.org/zap"
)

func TestCreateOrder(t *testing.T) {
    logger, _ := zap.NewDevelopment()
    repo := &Repository{}
    service := NewService(repo, logger)
    
    items := []OrderItem{
        {MenuItemID: "1", Name: "Pizza", Quantity: 2, Price: 15.99},
    }
    
    order, err := service.CreateOrder(context.Background(), "user123", "rest456", items)
    
    if err != nil {
        t.Errorf("CreateOrder failed: %v", err)
    }
    
    if order.ID == "" {
        t.Error("Order ID should not be empty")
    }
    
    if order.Status != "pending" {
        t.Errorf("Expected status 'pending', got '%s'", order.Status)
    }
    
    if order.Total != 31.98 {
        t.Errorf("Expected total 31.98, got %.2f", order.Total)
    }
}

func TestUpdateOrderStatus(t *testing.T) {
    logger, _ := zap.NewDevelopment()
    repo := &Repository{}
    service := NewService(repo, logger)
    
    order, err := service.UpdateOrderStatus(context.Background(), "order123", "confirmed")
    
    if err != nil {
        t.Errorf("UpdateOrderStatus failed: %v", err)
    }
    
    if order.Status != "confirmed" {
        t.Errorf("Expected status 'confirmed', got '%s'", order.Status)
    }
}

func TestCalculateTotal(t *testing.T) {
    items := []OrderItem{
        {Quantity: 2, Price: 10.0},
        {Quantity: 1, Price: 5.0},
    }
    
    total := calculateTotal(items)
    expected := 25.0
    
    if total != expected {
        t.Errorf("Expected total %.2f, got %.2f", expected, total)
    }
}

func TestGenerateOrderID(t *testing.T) {
    id1 := generateOrderID()
    id2 := generateOrderID()
    
    if id1 == "" {
        t.Error("Order ID should not be empty")
    }
    
    if id2 == "" {
        t.Error("Order ID should not be empty")
    }
    
    // IDs should be different
    if id1 == id2 {
        t.Errorf("Order IDs should be unique, got same: %s", id1)
    }
    
    // Check format
    if len(id1) < 10 {
        t.Errorf("Order ID format incorrect: %s", id1)
    }
}

func TestGetOrder(t *testing.T) {
    logger, _ := zap.NewDevelopment()
    repo := &Repository{}
    service := NewService(repo, logger)
    
    order, err := service.GetOrder(context.Background(), "order123")
    
    if err != nil {
        t.Errorf("GetOrder failed: %v", err)
    }
    
    if order.ID != "order123" {
        t.Errorf("Expected ID 'order123', got '%s'", order.ID)
    }
}

func TestCancelOrder(t *testing.T) {
    logger, _ := zap.NewDevelopment()
    repo := &Repository{}
    service := NewService(repo, logger)
    
    order, err := service.CancelOrder(context.Background(), "order123")
    
    if err != nil {
        t.Errorf("CancelOrder failed: %v", err)
    }
    
    if order.Status != "cancelled" {
        t.Errorf("Expected status 'cancelled', got '%s'", order.Status)
    }
}

func TestUpdateOrder(t *testing.T) {
    logger, _ := zap.NewDevelopment()
    repo := &Repository{}
    service := NewService(repo, logger)
    
    items := []OrderItem{
        {MenuItemID: "2", Name: "Burger", Quantity: 1, Price: 12.99},
    }
    
    order, err := service.UpdateOrder(context.Background(), "order123", items)
    
    if err != nil {
        t.Errorf("UpdateOrder failed: %v", err)
    }
    
    if order.ID != "order123" {
        t.Errorf("Expected ID 'order123', got '%s'", order.ID)
    }
}

func TestProcessPayment(t *testing.T) {
    logger, _ := zap.NewDevelopment()
    repo := &Repository{}
    service := NewService(repo, logger)
    
    paymentID, err := service.ProcessPayment(context.Background(), "order123", 99.99, "card")
    
    if err != nil {
        t.Errorf("ProcessPayment failed: %v", err)
    }
    
    if paymentID == "" {
        t.Error("Payment ID should not be empty")
    }
}

func TestTrackDelivery(t *testing.T) {
    logger, _ := zap.NewDevelopment()
    repo := &Repository{}
    service := NewService(repo, logger)
    
    delivery, err := service.TrackDelivery(context.Background(), "order123")
    
    if err != nil {
        t.Errorf("TrackDelivery failed: %v", err)
    }
    
    if delivery.Status == "" {
        t.Error("Delivery status should not be empty")
    }
    
    if delivery.CourierName == "" {
        t.Error("Courier name should not be empty")
    }
}

func TestApplyDiscount(t *testing.T) {
    logger, _ := zap.NewDevelopment()
    repo := &Repository{}
    service := NewService(repo, logger)
    
    order, err := service.ApplyDiscount(context.Background(), "order123", "DISCOUNT10")
    
    if err != nil {
        t.Errorf("ApplyDiscount failed: %v", err)
    }
    
    if order.Total != 90.0 {
        t.Errorf("Expected total 90.0, got %.2f", order.Total)
    }
}

// Benchmark tests
func BenchmarkCreateOrder(b *testing.B) {
    logger, _ := zap.NewDevelopment()
    repo := &Repository{}
    service := NewService(repo, logger)
    
    items := []OrderItem{
        {MenuItemID: "1", Name: "Pizza", Quantity: 2, Price: 15.99},
    }
    
    for i := 0; i < b.N; i++ {
        service.CreateOrder(context.Background(), "user123", "rest456", items)
    }
}

func BenchmarkGenerateOrderID(b *testing.B) {
    for i := 0; i < b.N; i++ {
        generateOrderID()
    }
}
