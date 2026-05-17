package payment

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type PaymentRepo struct {
	DB       *sql.DB
	NATS     *nats.Conn
	EmailSvc func(to, subject, body string) error
}

func NewRepository(db *sql.DB, nats *nats.Conn, emailSvc func(to, subject, body string) error) *PaymentRepo {
	return &PaymentRepo{DB: db, NATS: nats, EmailSvc: emailSvc}
}

func (r *PaymentRepo) CreatePayment(ctx context.Context, orderID, userID string, amount float64, method string) (*Payment, error) {
	id := uuid.NewString()
	status := "pending"
	_, err := r.DB.ExecContext(ctx, `INSERT INTO payments (id, order_id, user_id, amount, status, method, created_at) VALUES ($1, $2, $3, $4, $5, $6, NOW())`, id, orderID, userID, amount, status, method)
	if err != nil {
		return nil, err
	}
	return &Payment{ID: id, OrderID: orderID, UserID: userID, Amount: amount, Status: status, Method: method, CreatedAt: time.Now().UTC().Format(time.RFC3339)}, nil
}

func (r *PaymentRepo) GetPayment(ctx context.Context, paymentID string) (*Payment, error) {
	p := &Payment{}
	if err := r.DB.QueryRowContext(ctx, `SELECT id, order_id, user_id, amount, status, method, created_at FROM payments WHERE id = $1`, paymentID).Scan(&p.ID, &p.OrderID, &p.UserID, &p.Amount, &p.Status, &p.Method, &p.CreatedAt); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PaymentRepo) HandleOrderCreated(ctx context.Context, data []byte) error {
	var event OrderCreatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}
	status := "failed"
	if mockCharge(event.Total) {
		status = "success"
	}
	paymentID := uuid.NewString()
	_, err := r.DB.ExecContext(ctx, `INSERT INTO payments (id, order_id, user_id, amount, status, method, created_at) VALUES ($1, $2, $3, $4, $5, $6, NOW())`, paymentID, event.OrderID, event.UserID, event.Total, status, "card")
	if err != nil {
		return err
	}
	_, err = r.DB.ExecContext(ctx, `INSERT INTO transactions (id, payment_id, status, amount, created_at) VALUES ($1, $2, $3, $4, NOW())`, uuid.NewString(), paymentID, status, event.Total)
	if err != nil {
		return err
	}
	if err := r.publishPaymentProcessed(ctx, paymentID, event.OrderID, status); err != nil {
		return err
	}
	return nil
}

func (r *PaymentRepo) publishPaymentProcessed(ctx context.Context, paymentID, orderID, status string) error {
	event := PaymentProcessedEvent{PaymentID: paymentID, OrderID: orderID, Status: status}
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}
	if err := r.NATS.Publish("PaymentProcessed", payload); err != nil {
		return err
	}
	if r.EmailSvc != nil {
		body := fmt.Sprintf("Your payment for order %s has status %s.", orderID, status)
		_ = r.EmailSvc("customer@demo.com", "Payment Confirmation", body)
	}
	return nil
}

func mockCharge(amount float64) bool {
	return rand.Float64() > 0.1 && amount > 0
}
