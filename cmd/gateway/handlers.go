package main

import (
	"net/http"

	"food-delivery/internal/auth"
	"food-delivery/proto"
	"github.com/gin-gonic/gin"
)

func registerHandler(svc *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.RegisterRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, token, err := svc.Register(c.Request.Context(), req.Email, req.Password, req.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"token": token, "user": user})
	}
}

func loginHandler(svc *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.AuthRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, token, err := svc.Login(c.Request.Context(), req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
	}
}

func profileHandler(svc *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := getUserID(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		user, err := svc.GetUser(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func listRestaurantsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, []gin.H{
		{"id": "rest-1", "name": "Maple Kitchen", "cuisine": "American"},
		{"id": "rest-2", "name": "Saffron Grill", "cuisine": "Indian"},
	})
}

func getMenuHandler(c *gin.Context) {
	c.JSON(http.StatusOK, []gin.H{
		{"id": "item-1", "name": "Cheeseburger", "price": 12.5},
		{"id": "item-2", "name": "Masala Dosa", "price": 9.9},
	})
}

func addMenuItemHandler(c *gin.Context) {
	var payload struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": "menu-" + payload.Name, "name": payload.Name, "price": payload.Price})
}

func orderCreateHandler(client proto.OrderServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.CreateOrderRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID, err := getUserID(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		req.UserId = userID
		resp, err := client.CreateOrder(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, resp.Order)
	}
}

func orderListHandler(client proto.OrderServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := getUserID(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.ListUserOrders(c.Request.Context(), &proto.ListUserOrdersRequest{UserId: userID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Orders)
	}
}

func orderGetHandler(client proto.OrderServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := client.GetOrder(c.Request.Context(), &proto.GetOrderRequest{OrderId: c.Param("id")})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Order)
	}
}

func orderUpdateStatusHandler(client proto.OrderServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload struct {
			Status string `json:"status"`
		}
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.UpdateOrderStatus(c.Request.Context(), &proto.UpdateOrderStatusRequest{OrderId: c.Param("id"), Status: payload.Status})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Order)
	}
}

func orderCancelHandler(client proto.OrderServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := client.CancelOrder(c.Request.Context(), &proto.CancelOrderRequest{OrderId: c.Param("id")})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Order)
	}
}

func paymentCreateHandler(client proto.PaymentServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.CreatePaymentRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID, err := getUserID(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		req.UserId = userID
		resp, err := client.CreatePayment(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, resp.Payment)
	}
}

func paymentStatusHandler(client proto.PaymentServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := client.GetPaymentStatus(c.Request.Context(), &proto.GetPaymentStatusRequest{PaymentId: c.Param("id")})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Payment)
	}
}

func assignDeliveryHandler(client proto.DeliveryServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.AssignDeliveryRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.AssignDeliveryPerson(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, resp.Delivery)
	}
}

func updateDeliveryLocationHandler(client proto.DeliveryServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.UpdateDeliveryLocationRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.DeliveryId = c.Param("id")
		resp, err := client.UpdateDeliveryLocation(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Delivery)
	}
}

func deliveryStatusHandler(client proto.DeliveryServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := client.GetDeliveryStatus(c.Request.Context(), &proto.GetDeliveryStatusRequest{OrderId: c.Param("order_id")})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Delivery)
	}
}

func listDeliveriesHandler(client proto.DeliveryServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		personID := c.Query("person_id")
		resp, err := client.ListDeliveries(c.Request.Context(), &proto.ListDeliveriesRequest{DeliveryPersonId: personID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp.Deliveries)
	}
}
