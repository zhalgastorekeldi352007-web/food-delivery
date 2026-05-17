package delivery

type Delivery struct {
	ID               string `json:"id"`
	OrderID          string `json:"order_id"`
	DeliveryPersonID string `json:"delivery_person_id"`
	Status           string `json:"status"`
	CurrentLocation  string `json:"current_location"`
	AssignedAt       string `json:"assigned_at"`
}

type DeliveryPerson struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Status string `json:"status"`
}

type PaymentProcessedEvent struct {
	PaymentID string `json:"payment_id"`
	OrderID   string `json:"order_id"`
	Status    string `json:"status"`
}

type DeliveryAssignedEvent struct {
	DeliveryID       string `json:"delivery_id"`
	OrderID          string `json:"order_id"`
	DeliveryPersonID string `json:"delivery_person_id"`
}
