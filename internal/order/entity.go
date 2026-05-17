package order

import (
    "fmt"
    "time"
    "math/rand"
)

type OrderItem struct {
    MenuItemID string
    Name       string
    Quantity   int32
    Price      float64
}

type Order struct {
    ID           string
    UserID       string
    RestaurantID string
    Items        []OrderItem
    Status       string
    Total        float64
    CreatedAt    time.Time
}

type CreateOrderRequest struct {
    UserID       string
    RestaurantID string
    Items        []OrderItem
}

// Глобальный счетчик для уникальности ID
var idCounter = 0

func generateOrderID() string {
    idCounter++
    return fmt.Sprintf("ORD-%d-%d", time.Now().UnixNano(), idCounter)
}

func calculateTotal(items []OrderItem) float64 {
    var total float64
    for _, item := range items {
        total += float64(item.Quantity) * item.Price
    }
    return total
}

func init() {
    rand.Seed(time.Now().UnixNano())
}
