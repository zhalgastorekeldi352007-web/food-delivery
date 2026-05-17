package delivery

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type Repository struct {
	DB       *sql.DB
	NATS     *nats.Conn
	EmailSvc func(to, subject, body string) error
}

func NewRepository(db *sql.DB, nats *nats.Conn, emailSvc func(to, subject, body string) error) *Repository {
	return &Repository{DB: db, NATS: nats, EmailSvc: emailSvc}
}

func (r *Repository) AssignDeliveryPerson(ctx context.Context, orderID, deliveryPersonID string) (*Delivery, error) {
	id := uuid.NewString()
	assignedAt := time.Now().UTC().Format(time.RFC3339)
	status := "assigned"
	_, err := r.DB.ExecContext(ctx, `INSERT INTO deliveries (id, order_id, delivery_person_id, status, current_location, assigned_at) VALUES ($1, $2, $3, $4, $5, $6)`, id, orderID, deliveryPersonID, status, "warehouse", assignedAt)
	if err != nil {
		return nil, err
	}
	if err := r.publishDeliveryAssigned(ctx, id, orderID, deliveryPersonID); err != nil {
		return nil, err
	}
	if r.EmailSvc != nil {
		body := fmt.Sprintf("Delivery person %s has been assigned to order %s.", deliveryPersonID, orderID)
		_ = r.EmailSvc("delivery@demo.com", "Delivery Assigned", body)
	}
	return &Delivery{ID: id, OrderID: orderID, DeliveryPersonID: deliveryPersonID, Status: status, CurrentLocation: "warehouse", AssignedAt: assignedAt}, nil
}

func (r *Repository) UpdateDeliveryLocation(ctx context.Context, deliveryID, location string) (*Delivery, error) {
	_, err := r.DB.ExecContext(ctx, `UPDATE deliveries SET current_location = $1 WHERE id = $2`, location, deliveryID)
	if err != nil {
		return nil, err
	}
	return r.GetDeliveryByID(ctx, deliveryID)
}

func (r *Repository) GetDeliveryStatus(ctx context.Context, orderID string) (*Delivery, error) {
	delivery := &Delivery{}
	if err := r.DB.QueryRowContext(ctx, `SELECT id, order_id, delivery_person_id, status, current_location, assigned_at FROM deliveries WHERE order_id = $1`, orderID).Scan(&delivery.ID, &delivery.OrderID, &delivery.DeliveryPersonID, &delivery.Status, &delivery.CurrentLocation, &delivery.AssignedAt); err != nil {
		return nil, err
	}
	return delivery, nil
}

func (r *Repository) ListDeliveries(ctx context.Context, deliveryPersonID string) ([]*Delivery, error) {
	rows, err := r.DB.QueryContext(ctx, `SELECT id, order_id, delivery_person_id, status, current_location, assigned_at FROM deliveries WHERE delivery_person_id = $1`, deliveryPersonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var deliveries []*Delivery
	for rows.Next() {
		d := &Delivery{}
		if err := rows.Scan(&d.ID, &d.OrderID, &d.DeliveryPersonID, &d.Status, &d.CurrentLocation, &d.AssignedAt); err != nil {
			return nil, err
		}
		deliveries = append(deliveries, d)
	}
	return deliveries, nil
}

func (r *Repository) HandlePaymentProcessed(ctx context.Context, data []byte) error {
	var event PaymentProcessedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}
	if event.Status != "success" {
		return nil
	}
	return nil
}

func (r *Repository) publishDeliveryAssigned(ctx context.Context, deliveryID, orderID, deliveryPersonID string) error {
	event := DeliveryAssignedEvent{DeliveryID: deliveryID, OrderID: orderID, DeliveryPersonID: deliveryPersonID}
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return r.NATS.Publish("DeliveryAssigned", payload)
}

func (r *Repository) GetDeliveryByID(ctx context.Context, deliveryID string) (*Delivery, error) {
	d := &Delivery{}
	if err := r.DB.QueryRowContext(ctx, `SELECT id, order_id, delivery_person_id, status, current_location, assigned_at FROM deliveries WHERE id = $1`, deliveryID).Scan(&d.ID, &d.OrderID, &d.DeliveryPersonID, &d.Status, &d.CurrentLocation, &d.AssignedAt); err != nil {
		return nil, err
	}
	return d, nil
}
