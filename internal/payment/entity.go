package payment

type Payment struct {
	ID        string  `json:"id"`
	OrderID   string  `json:"order_id"`
	UserID    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	Method    string  `json:"method"`
	CreatedAt string  `json:"created_at"`
}

type Transaction struct {
	ID        string  `json:"id"`
	PaymentID string  `json:"payment_id"`
	Status    string  `json:"status"`
	Amount    float64 `json:"amount"`
}

type OrderCreatedEvent struct {
	OrderID string  `json:"order_id"`
	UserID  string  `json:"user_id"`
	Total   float64 `json:"total"`
}

type PaymentProcessedEvent struct {
	PaymentID string `json:"payment_id"`
	OrderID   string `json:"order_id"`
	Status    string `json:"status"`
}
