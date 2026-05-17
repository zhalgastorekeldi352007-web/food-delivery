package main

import (
    "context"
    "fmt"
    "time"

    "food-delivery/internal/order"
    "food-delivery/proto"
    "go.uber.org/zap"
)

type orderServer struct {
    service *order.Service
    log     *zap.Logger
}

// Existing endpoints (5)
func (s *orderServer) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.OrderResponse, error) {
    items := make([]order.OrderItem, 0, len(req.Items))
    for _, item := range req.Items {
        items = append(items, order.OrderItem{
            MenuItemID: item.MenuItemId,
            Name:       item.Name,
            Quantity:   item.Quantity,
            Price:      item.Price,
        })
    }
    orderEntity, err := s.service.CreateOrder(ctx, req.UserId, req.RestaurantId, items)
    if err != nil {
        return nil, err
    }
    return &proto.OrderResponse{Order: toProtoOrder(orderEntity)}, nil
}

func (s *orderServer) UpdateOrderStatus(ctx context.Context, req *proto.UpdateOrderStatusRequest) (*proto.OrderResponse, error) {
    orderEntity, err := s.service.UpdateOrderStatus(ctx, req.OrderId, req.Status)
    if err != nil {
        return nil, err
    }
    return &proto.OrderResponse{Order: toProtoOrder(orderEntity)}, nil
}

func (s *orderServer) GetOrder(ctx context.Context, req *proto.GetOrderRequest) (*proto.OrderResponse, error) {
    orderEntity, err := s.service.GetOrder(ctx, req.OrderId)
    if err != nil {
        return nil, err
    }
    return &proto.OrderResponse{Order: toProtoOrder(orderEntity)}, nil
}

func (s *orderServer) ListUserOrders(ctx context.Context, req *proto.ListUserOrdersRequest) (*proto.ListOrdersResponse, error) {
    orders, err := s.service.ListUserOrders(ctx, req.UserId)
    if err != nil {
        return nil, err
    }
    result := &proto.ListOrdersResponse{}
    for _, orderEntity := range orders {
        result.Orders = append(result.Orders, toProtoOrder(orderEntity))
    }
    return result, nil
}

func (s *orderServer) CancelOrder(ctx context.Context, req *proto.CancelOrderRequest) (*proto.OrderResponse, error) {
    orderEntity, err := s.service.CancelOrder(ctx, req.OrderId)
    if err != nil {
        return nil, err
    }
    return &proto.OrderResponse{Order: toProtoOrder(orderEntity)}, nil
}

// New endpoints (7 more to reach 12+)

// 6. UpdateOrder - Update entire order
func (s *orderServer) UpdateOrder(ctx context.Context, req *proto.UpdateOrderRequest) (*proto.OrderResponse, error) {
    s.log.Info("UpdateOrder called", zap.String("order_id", req.OrderId))
    
    items := make([]order.OrderItem, 0, len(req.Items))
    for _, item := range req.Items {
        items = append(items, order.OrderItem{
            MenuItemID: item.MenuItemId,
            Name:       item.Name,
            Quantity:   item.Quantity,
            Price:      item.Price,
        })
    }
    
    orderEntity, err := s.service.UpdateOrder(ctx, req.OrderId, items)
    if err != nil {
        return nil, err
    }
    return &proto.OrderResponse{Order: toProtoOrder(orderEntity)}, nil
}

// 7. DeleteOrder - Delete/cancel order
func (s *orderServer) DeleteOrder(ctx context.Context, req *proto.DeleteOrderRequest) (*proto.DeleteOrderResponse, error) {
    s.log.Info("DeleteOrder called", zap.String("order_id", req.OrderId))
    
    err := s.service.DeleteOrder(ctx, req.OrderId)
    if err != nil {
        return &proto.DeleteOrderResponse{Success: false, Message: err.Error()}, err
    }
    return &proto.DeleteOrderResponse{Success: true, Message: "Order deleted successfully"}, nil
}

// 8. ProcessPayment - Process payment for order
func (s *orderServer) ProcessPayment(ctx context.Context, req *proto.ProcessPaymentRequest) (*proto.PaymentResponse, error) {
    s.log.Info("ProcessPayment called", zap.String("order_id", req.OrderId), zap.Float64("amount", req.Amount))
    
    paymentID, err := s.service.ProcessPayment(ctx, req.OrderId, req.Amount, req.PaymentMethod)
    if err != nil {
        return &proto.PaymentResponse{Status: "failed", Message: err.Error()}, err
    }
    return &proto.PaymentResponse{
        PaymentId: paymentID,
        Status:    "success",
        Message:   "Payment processed successfully",
    }, nil
}

// 9. GetPaymentStatus - Check payment status
func (s *orderServer) GetPaymentStatus(ctx context.Context, req *proto.GetPaymentStatusRequest) (*proto.PaymentStatusResponse, error) {
    s.log.Info("GetPaymentStatus called", zap.String("order_id", req.OrderId))
    
    status, amount, err := s.service.GetPaymentStatus(ctx, req.OrderId)
    if err != nil {
        return nil, err
    }
    return &proto.PaymentStatusResponse{
        OrderId: req.OrderId,
        Status:  status,
        Amount:  amount,
    }, nil
}

// 10. TrackDelivery - Track delivery status
func (s *orderServer) TrackDelivery(ctx context.Context, req *proto.TrackDeliveryRequest) (*proto.DeliveryStatusResponse, error) {
    s.log.Info("TrackDelivery called", zap.String("order_id", req.OrderId))
    
    delivery, err := s.service.TrackDelivery(ctx, req.OrderId)
    if err != nil {
        return nil, err
    }
    return &proto.DeliveryStatusResponse{
        OrderId:          req.OrderId,
        Status:           delivery.Status,
        CourierName:      delivery.CourierName,
        CourierPhone:     delivery.CourierPhone,
        EstimatedMinutes: delivery.EstimatedMinutes,
    }, nil
}

// 11. UpdateDeliveryStatus - Update delivery status (internal)
func (s *orderServer) UpdateDeliveryStatus(ctx context.Context, req *proto.UpdateDeliveryStatusRequest) (*proto.DeliveryStatusResponse, error) {
    s.log.Info("UpdateDeliveryStatus called", 
        zap.String("order_id", req.OrderId),
        zap.String("status", req.Status),
        zap.String("location", req.Location))
    
    err := s.service.UpdateDeliveryStatus(ctx, req.OrderId, req.Status, req.Location)
    if err != nil {
        return nil, err
    }
    return &proto.DeliveryStatusResponse{
        OrderId: req.OrderId,
        Status:  req.Status,
    }, nil
}

// 12. ApplyDiscount - Apply discount code to order
func (s *orderServer) ApplyDiscount(ctx context.Context, req *proto.ApplyDiscountRequest) (*proto.OrderResponse, error) {
    s.log.Info("ApplyDiscount called", 
        zap.String("order_id", req.OrderId),
        zap.String("discount_code", req.DiscountCode))
    
    orderEntity, err := s.service.ApplyDiscount(ctx, req.OrderId, req.DiscountCode)
    if err != nil {
        return nil, err
    }
    return &proto.OrderResponse{Order: toProtoOrder(orderEntity)}, nil
}

// 13. GetOrderHistory - Get user order history
func (s *orderServer) GetOrderHistory(ctx context.Context, req *proto.GetOrderHistoryRequest) (*proto.OrderHistoryResponse, error) {
    s.log.Info("GetOrderHistory called", zap.String("user_id", req.UserId), zap.Int32("limit", req.Limit))
    
    orders, total, spent, err := s.service.GetOrderHistory(ctx, req.UserId, req.Limit)
    if err != nil {
        return nil, err
    }
    
    protoOrders := make([]*proto.Order, 0, len(orders))
    for _, o := range orders {
        protoOrders = append(protoOrders, toProtoOrder(o))
    }
    
    return &proto.OrderHistoryResponse{
        Orders:      protoOrders,
        TotalOrders: total,
        TotalSpent:  spent,
    }, nil
}

func toProtoOrder(orderEntity *order.Order) *proto.Order {
    items := make([]*proto.OrderItem, 0, len(orderEntity.Items))
    for _, item := range orderEntity.Items {
        items = append(items, &proto.OrderItem{
            MenuItemId: item.MenuItemID,
            Name:       item.Name,
            Quantity:   item.Quantity,
            Price:      item.Price,
        })
    }
    return &proto.Order{
        Id:           orderEntity.ID,
        UserId:       orderEntity.UserID,
        RestaurantId: orderEntity.RestaurantID,
        Items:        items,
        Status:       orderEntity.Status,
        Total:        orderEntity.Total,
        CreatedAt:    orderEntity.CreatedAt.Format(time.RFC3339),
    }
}
