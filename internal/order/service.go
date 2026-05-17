package order

import (
    "context"
    "time"

    "food-delivery/internal/email"
    "go.uber.org/zap"
)

type Service struct {
    repo         *Repository
    log          *zap.Logger
    emailService *email.EmailService
}

func NewService(repo *Repository, log *zap.Logger) *Service {
    return &Service{
        repo:         repo,
        log:          log,
        emailService: email.NewEmailService(),
    }
}

func (s *Service) CreateOrder(ctx context.Context, userID, restaurantID string, items []OrderItem) (*Order, error) {
    s.log.Info("Creating order", zap.String("userID", userID), zap.String("restaurantID", restaurantID))

    order := &Order{
        ID:           generateOrderID(),
        UserID:       userID,
        RestaurantID: restaurantID,
        Items:        items,
        Status:       "pending",
        Total:        calculateTotal(items),
        CreatedAt:    time.Now(),
    }

    // Send email confirmation
    s.sendOrderConfirmation(order, userID+"@example.com")

    return order, nil
}

func (s *Service) UpdateOrderStatus(ctx context.Context, orderID, status string) (*Order, error) {
    s.log.Info("Updating order status", zap.String("orderID", orderID), zap.String("status", status))

    order := &Order{
        ID:     orderID,
        Status: status,
    }

    // Send status update email
    s.sendOrderStatusUpdate(order, orderID+"@example.com", status)

    return order, nil
}

func (s *Service) GetOrder(ctx context.Context, orderID string) (*Order, error) {
    s.log.Info("Getting order", zap.String("orderID", orderID))

    order := &Order{
        ID:     orderID,
        Status: "pending",
    }

    return order, nil
}

func (s *Service) ListUserOrders(ctx context.Context, userID string) ([]*Order, error) {
    s.log.Info("Listing user orders", zap.String("userID", userID))

    orders := []*Order{}

    return orders, nil
}

func (s *Service) CancelOrder(ctx context.Context, orderID string) (*Order, error) {
    s.log.Info("Cancelling order", zap.String("orderID", orderID))

    order := &Order{
        ID:     orderID,
        Status: "cancelled",
    }

    return order, nil
}

func (s *Service) UpdateOrder(ctx context.Context, orderID string, items []OrderItem) (*Order, error) {
    s.log.Info("Updating order", zap.String("order_id", orderID))
    return &Order{ID: orderID, Items: items}, nil
}

func (s *Service) DeleteOrder(ctx context.Context, orderID string) error {
    s.log.Info("Deleting order", zap.String("order_id", orderID))
    return nil
}

func (s *Service) ProcessPayment(ctx context.Context, orderID string, amount float64, method string) (string, error) {
    s.log.Info("Processing payment", zap.String("order_id", orderID), zap.Float64("amount", amount))
    return "pay_" + orderID, nil
}

func (s *Service) GetPaymentStatus(ctx context.Context, orderID string) (string, float64, error) {
    return "completed", 100.0, nil
}

func (s *Service) TrackDelivery(ctx context.Context, orderID string) (*DeliveryInfo, error) {
    return &DeliveryInfo{
        Status:           "on_the_way",
        CourierName:      "John Doe",
        CourierPhone:     "+123456789",
        EstimatedMinutes: 30,
    }, nil
}

func (s *Service) UpdateDeliveryStatus(ctx context.Context, orderID, status, location string) error {
    return nil
}

func (s *Service) ApplyDiscount(ctx context.Context, orderID, discountCode string) (*Order, error) {
    return &Order{ID: orderID, Total: 90.0}, nil
}

func (s *Service) GetOrderHistory(ctx context.Context, userID string, limit int32) ([]*Order, int32, float64, error) {
    return []*Order{}, 0, 0.0, nil
}

func (s *Service) sendOrderConfirmation(order *Order, userEmail string) {
    go func() {
        if err := s.emailService.SendOrderConfirmation(userEmail, order.ID, order.Total); err != nil {
            s.log.Error("Failed to send email", zap.Error(err))
        } else {
            s.log.Info("Order confirmation email sent", zap.String("order_id", order.ID))
        }
    }()
}

func (s *Service) sendOrderStatusUpdate(order *Order, userEmail string, status string) {
    go func() {
        if err := s.emailService.SendOrderStatusUpdate(userEmail, order.ID, status); err != nil {
            s.log.Error("Failed to send status update email", zap.Error(err))
        } else {
            s.log.Info("Order status update email sent", zap.String("order_id", order.ID), zap.String("status", status))
        }
    }()
}

type DeliveryInfo struct {
    Status           string
    CourierName      string
    CourierPhone     string
    EstimatedMinutes int32
}
