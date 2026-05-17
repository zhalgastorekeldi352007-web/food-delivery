package proto

import (
	"context"

	"google.golang.org/grpc"
)

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserProfileRequest struct {
	UserId string `json:"user_id"`
}

type UserProfileResponse struct {
	User *User `json:"user"`
}

type OrderItem struct {
	MenuItemId string  `json:"menu_item_id"`
	Name       string  `json:"name"`
	Quantity   int32   `json:"quantity"`
	Price      float64 `json:"price"`
}

type Order struct {
	Id           string      `json:"id"`
	UserId       string      `json:"user_id"`
	RestaurantId string      `json:"restaurant_id"`
	Items        []OrderItem `json:"items"`
	Status       string      `json:"status"`
	Total        float64     `json:"total"`
	CreatedAt    string      `json:"created_at"`
}

type CreateOrderRequest struct {
	UserId       string      `json:"user_id"`
	RestaurantId string      `json:"restaurant_id"`
	Items        []OrderItem `json:"items"`
}

type UpdateOrderStatusRequest struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
}

type GetOrderRequest struct {
	OrderId string `json:"order_id"`
}

type ListUserOrdersRequest struct {
	UserId string `json:"user_id"`
}

type CancelOrderRequest struct {
	OrderId string `json:"order_id"`
}

type OrderResponse struct {
	Order *Order `json:"order"`
}

type ListOrdersResponse struct {
	Orders []*Order `json:"orders"`
}

type Payment struct {
	Id        string  `json:"id"`
	OrderId   string  `json:"order_id"`
	UserId    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	Method    string  `json:"method"`
	CreatedAt string  `json:"created_at"`
}

type CreatePaymentRequest struct {
	OrderId string  `json:"order_id"`
	UserId  string  `json:"user_id"`
	Amount  float64 `json:"amount"`
	Method  string  `json:"method"`
}

type GetPaymentStatusRequest struct {
	PaymentId string `json:"payment_id"`
}

type PaymentResponse struct {
	Payment *Payment `json:"payment"`
}

type Delivery struct {
	Id               string `json:"id"`
	OrderId          string `json:"order_id"`
	DeliveryPersonId string `json:"delivery_person_id"`
	Status           string `json:"status"`
	CurrentLocation  string `json:"current_location"`
	AssignedAt       string `json:"assigned_at"`
}

type DeliveryPerson struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Status string `json:"status"`
}

type AssignDeliveryRequest struct {
	OrderId          string `json:"order_id"`
	DeliveryPersonId string `json:"delivery_person_id"`
}

type UpdateDeliveryLocationRequest struct {
	DeliveryId string `json:"delivery_id"`
	Location   string `json:"location"`
}

type GetDeliveryStatusRequest struct {
	OrderId string `json:"order_id"`
}

type ListDeliveriesRequest struct {
	DeliveryPersonId string `json:"delivery_person_id"`
}

type DeliveryResponse struct {
	Delivery *Delivery `json:"delivery"`
}

type ListDeliveriesResponse struct {
	Deliveries []*Delivery `json:"deliveries"`
}

// AuthService

type AuthServiceServer interface {
	Register(context.Context, *RegisterRequest) (*AuthResponse, error)
	Login(context.Context, *AuthRequest) (*AuthResponse, error)
	GetProfile(context.Context, *UserProfileRequest) (*UserProfileResponse, error)
}

type AuthServiceClient interface {
	Register(context.Context, *RegisterRequest, ...grpc.CallOption) (*AuthResponse, error)
	Login(context.Context, *AuthRequest, ...grpc.CallOption) (*AuthResponse, error)
	GetProfile(context.Context, *UserProfileRequest, ...grpc.CallOption) (*UserProfileResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.AuthService/Register", in, out, opts...)
	return out, err
}

func (c *authServiceClient) Login(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.AuthService/Login", in, out, opts...)
	return out, err
}

func (c *authServiceClient) GetProfile(ctx context.Context, in *UserProfileRequest, opts ...grpc.CallOption) (*UserProfileResponse, error) {
	out := new(UserProfileResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.AuthService/GetProfile", in, out, opts...)
	return out, err
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&grpc.ServiceDesc{
		ServiceName: "fooddelivery.AuthService",
		HandlerType: (*AuthServiceServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Register",
				Handler:    authRegisterHandler,
			},
			{
				MethodName: "Login",
				Handler:    authLoginHandler,
			},
			{
				MethodName: "GetProfile",
				Handler:    authGetProfileHandler,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "food.proto",
	}, srv)
}

func authRegisterHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.AuthService/Register"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func authLoginHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.AuthService/Login"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func authGetProfileHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProfileRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.AuthService/GetProfile"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetProfile(ctx, req.(*UserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService

type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*OrderResponse, error)
	UpdateOrderStatus(context.Context, *UpdateOrderStatusRequest) (*OrderResponse, error)
	GetOrder(context.Context, *GetOrderRequest) (*OrderResponse, error)
	ListUserOrders(context.Context, *ListUserOrdersRequest) (*ListOrdersResponse, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*OrderResponse, error)
}

type OrderServiceClient interface {
	CreateOrder(context.Context, *CreateOrderRequest, ...grpc.CallOption) (*OrderResponse, error)
	UpdateOrderStatus(context.Context, *UpdateOrderStatusRequest, ...grpc.CallOption) (*OrderResponse, error)
	GetOrder(context.Context, *GetOrderRequest, ...grpc.CallOption) (*OrderResponse, error)
	ListUserOrders(context.Context, *ListUserOrdersRequest, ...grpc.CallOption) (*ListOrdersResponse, error)
	CancelOrder(context.Context, *CancelOrderRequest, ...grpc.CallOption) (*OrderResponse, error)
}

type orderServiceClient struct{ cc *grpc.ClientConn }

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient { return &orderServiceClient{cc} }

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.OrderService/CreateOrder", in, out, opts...)
	return out, err
}

func (c *orderServiceClient) UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.OrderService/UpdateOrderStatus", in, out, opts...)
	return out, err
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.OrderService/GetOrder", in, out, opts...)
	return out, err
}

func (c *orderServiceClient) ListUserOrders(ctx context.Context, in *ListUserOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error) {
	out := new(ListOrdersResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.OrderService/ListUserOrders", in, out, opts...)
	return out, err
}

func (c *orderServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.OrderService/CancelOrder", in, out, opts...)
	return out, err
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&grpc.ServiceDesc{
		ServiceName: "fooddelivery.OrderService",
		HandlerType: (*OrderServiceServer)(nil),
		Methods: []grpc.MethodDesc{
			{MethodName: "CreateOrder", Handler: orderCreateOrderHandler},
			{MethodName: "UpdateOrderStatus", Handler: orderUpdateOrderStatusHandler},
			{MethodName: "GetOrder", Handler: orderGetOrderHandler},
			{MethodName: "ListUserOrders", Handler: orderListUserOrdersHandler},
			{MethodName: "CancelOrder", Handler: orderCancelOrderHandler},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "food.proto",
	}, srv)
}

func orderCreateOrderHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.OrderService/CreateOrder"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func orderUpdateOrderStatusHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderStatusRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.OrderService/UpdateOrderStatus"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrderStatus(ctx, req.(*UpdateOrderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func orderGetOrderHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.OrderService/GetOrder"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func orderListUserOrdersHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserOrdersRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListUserOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.OrderService/ListUserOrders"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListUserOrders(ctx, req.(*ListUserOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func orderCancelOrderHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.OrderService/CancelOrder"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService

type PaymentServiceServer interface {
	CreatePayment(context.Context, *CreatePaymentRequest) (*PaymentResponse, error)
	GetPaymentStatus(context.Context, *GetPaymentStatusRequest) (*PaymentResponse, error)
}

type PaymentServiceClient interface {
	CreatePayment(context.Context, *CreatePaymentRequest, ...grpc.CallOption) (*PaymentResponse, error)
	GetPaymentStatus(context.Context, *GetPaymentStatusRequest, ...grpc.CallOption) (*PaymentResponse, error)
}

type paymentServiceClient struct{ cc *grpc.ClientConn }

func NewPaymentServiceClient(cc *grpc.ClientConn) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.PaymentService/CreatePayment", in, out, opts...)
	return out, err
}

func (c *paymentServiceClient) GetPaymentStatus(ctx context.Context, in *GetPaymentStatusRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.PaymentService/GetPaymentStatus", in, out, opts...)
	return out, err
}

func RegisterPaymentServiceServer(s *grpc.Server, srv PaymentServiceServer) {
	s.RegisterService(&grpc.ServiceDesc{
		ServiceName: "fooddelivery.PaymentService",
		HandlerType: (*PaymentServiceServer)(nil),
		Methods: []grpc.MethodDesc{
			{MethodName: "CreatePayment", Handler: paymentCreatePaymentHandler},
			{MethodName: "GetPaymentStatus", Handler: paymentGetPaymentStatusHandler},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "food.proto",
	}, srv)
}

func paymentCreatePaymentHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.PaymentService/CreatePayment"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func paymentGetPaymentStatusHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentStatusRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPaymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.PaymentService/GetPaymentStatus"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPaymentStatus(ctx, req.(*GetPaymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeliveryService

type DeliveryServiceServer interface {
	AssignDeliveryPerson(context.Context, *AssignDeliveryRequest) (*DeliveryResponse, error)
	UpdateDeliveryLocation(context.Context, *UpdateDeliveryLocationRequest) (*DeliveryResponse, error)
	GetDeliveryStatus(context.Context, *GetDeliveryStatusRequest) (*DeliveryResponse, error)
	ListDeliveries(context.Context, *ListDeliveriesRequest) (*ListDeliveriesResponse, error)
}

type DeliveryServiceClient interface {
	AssignDeliveryPerson(context.Context, *AssignDeliveryRequest, ...grpc.CallOption) (*DeliveryResponse, error)
	UpdateDeliveryLocation(context.Context, *UpdateDeliveryLocationRequest, ...grpc.CallOption) (*DeliveryResponse, error)
	GetDeliveryStatus(context.Context, *GetDeliveryStatusRequest, ...grpc.CallOption) (*DeliveryResponse, error)
	ListDeliveries(context.Context, *ListDeliveriesRequest, ...grpc.CallOption) (*ListDeliveriesResponse, error)
}

type deliveryServiceClient struct{ cc *grpc.ClientConn }

func NewDeliveryServiceClient(cc *grpc.ClientConn) DeliveryServiceClient {
	return &deliveryServiceClient{cc}
}

func (c *deliveryServiceClient) AssignDeliveryPerson(ctx context.Context, in *AssignDeliveryRequest, opts ...grpc.CallOption) (*DeliveryResponse, error) {
	out := new(DeliveryResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.DeliveryService/AssignDeliveryPerson", in, out, opts...)
	return out, err
}

func (c *deliveryServiceClient) UpdateDeliveryLocation(ctx context.Context, in *UpdateDeliveryLocationRequest, opts ...grpc.CallOption) (*DeliveryResponse, error) {
	out := new(DeliveryResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.DeliveryService/UpdateDeliveryLocation", in, out, opts...)
	return out, err
}

func (c *deliveryServiceClient) GetDeliveryStatus(ctx context.Context, in *GetDeliveryStatusRequest, opts ...grpc.CallOption) (*DeliveryResponse, error) {
	out := new(DeliveryResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.DeliveryService/GetDeliveryStatus", in, out, opts...)
	return out, err
}

func (c *deliveryServiceClient) ListDeliveries(ctx context.Context, in *ListDeliveriesRequest, opts ...grpc.CallOption) (*ListDeliveriesResponse, error) {
	out := new(ListDeliveriesResponse)
	err := c.cc.Invoke(ctx, "/fooddelivery.DeliveryService/ListDeliveries", in, out, opts...)
	return out, err
}

func RegisterDeliveryServiceServer(s *grpc.Server, srv DeliveryServiceServer) {
	s.RegisterService(&grpc.ServiceDesc{
		ServiceName: "fooddelivery.DeliveryService",
		HandlerType: (*DeliveryServiceServer)(nil),
		Methods: []grpc.MethodDesc{
			{MethodName: "AssignDeliveryPerson", Handler: deliveryAssignHandler},
			{MethodName: "UpdateDeliveryLocation", Handler: deliveryUpdateLocationHandler},
			{MethodName: "GetDeliveryStatus", Handler: deliveryGetStatusHandler},
			{MethodName: "ListDeliveries", Handler: deliveryListHandler},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "food.proto",
	}, srv)
}

func deliveryAssignHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignDeliveryRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).AssignDeliveryPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.DeliveryService/AssignDeliveryPerson"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).AssignDeliveryPerson(ctx, req.(*AssignDeliveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func deliveryUpdateLocationHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeliveryLocationRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).UpdateDeliveryLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.DeliveryService/UpdateDeliveryLocation"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).UpdateDeliveryLocation(ctx, req.(*UpdateDeliveryLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func deliveryGetStatusHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeliveryStatusRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).GetDeliveryStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.DeliveryService/GetDeliveryStatus"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).GetDeliveryStatus(ctx, req.(*GetDeliveryStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func deliveryListHandler(srv interface{}, ctx context.Context, decodeFunc func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeliveriesRequest)
	if err := decodeFunc(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).ListDeliveries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/fooddelivery.DeliveryService/ListDeliveries"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).ListDeliveries(ctx, req.(*ListDeliveriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}
