package proto

type Order struct {
    Id           string
    UserId       string
    RestaurantId string
    Items        []*OrderItem
    Status       string
    Total        float64
    CreatedAt    string
}

type OrderItem struct {
    MenuItemId string
    Name       string
    Quantity   int32
    Price      float64
}

type OrderResponse struct {
    Order *Order
}

type ListOrdersResponse struct {
    Orders []*Order
    Total  int32
}

type GetOrderRequest struct {
    OrderId string
}

type CreateOrderRequest struct {
    UserId       string
    RestaurantId string
    Items        []*OrderItem
}

type UpdateOrderRequest struct {
    OrderId string
    Items   []*OrderItem
}

type DeleteOrderRequest struct {
    OrderId string
}

type DeleteOrderResponse struct {
    Success bool
    Message string
}

type ListOrdersRequest struct {
    UserId   string
    Page     int32
    PageSize int32
}

type UpdateOrderStatusRequest struct {
    OrderId string
    Status  string
}

type CancelOrderRequest struct {
    OrderId string
    Reason  string
}

type ProcessPaymentRequest struct {
    OrderId       string
    Amount        float64
    PaymentMethod string
}

type PaymentResponse struct {
    PaymentId string
    Status    string
    Message   string
}

type GetPaymentStatusRequest struct {
    OrderId string
}

type PaymentStatusResponse struct {
    OrderId string
    Status  string
    Amount  float64
}

type TrackDeliveryRequest struct {
    OrderId string
}

type DeliveryStatusResponse struct {
    OrderId          string
    Status           string
    CourierName      string
    CourierPhone     string
    EstimatedMinutes int32
}

type UpdateDeliveryStatusRequest struct {
    OrderId  string
    Status   string
    Location string
}

type ApplyDiscountRequest struct {
    OrderId      string
    DiscountCode string
}

type GetOrderHistoryRequest struct {
    UserId string
    Limit  int32
}

type OrderHistoryResponse struct {
    Orders      []*Order
    TotalOrders int32
    TotalSpent  float64
}

type OrderServiceServer interface {
    CreateOrder(ctx interface{}, req *CreateOrderRequest) (*OrderResponse, error)
    GetOrder(ctx interface{}, req *GetOrderRequest) (*OrderResponse, error)
    UpdateOrder(ctx interface{}, req *UpdateOrderRequest) (*OrderResponse, error)
    DeleteOrder(ctx interface{}, req *DeleteOrderRequest) (*DeleteOrderResponse, error)
    ListOrders(ctx interface{}, req *ListOrdersRequest) (*ListOrdersResponse, error)
    UpdateOrderStatus(ctx interface{}, req *UpdateOrderStatusRequest) (*OrderResponse, error)
    CancelOrder(ctx interface{}, req *CancelOrderRequest) (*OrderResponse, error)
    ProcessPayment(ctx interface{}, req *ProcessPaymentRequest) (*PaymentResponse, error)
    GetPaymentStatus(ctx interface{}, req *GetPaymentStatusRequest) (*PaymentStatusResponse, error)
    TrackDelivery(ctx interface{}, req *TrackDeliveryRequest) (*DeliveryStatusResponse, error)
    UpdateDeliveryStatus(ctx interface{}, req *UpdateDeliveryStatusRequest) (*DeliveryStatusResponse, error)
    ApplyDiscount(ctx interface{}, req *ApplyDiscountRequest) (*OrderResponse, error)
    GetOrderHistory(ctx interface{}, req *GetOrderHistoryRequest) (*OrderHistoryResponse, error)
}

func RegisterOrderServiceServer(s interface{}, srv OrderServiceServer) {}
