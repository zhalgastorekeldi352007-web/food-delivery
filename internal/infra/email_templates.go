package infra

const (
	OrderConfirmationTemplate = `<!DOCTYPE html>
<html><body><h1>Order Confirmation</h1><p>Your order {{.OrderID}} is confirmed and being processed.</p></body></html>`
	PaymentConfirmationTemplate = `<!DOCTYPE html>
<html><body><h1>Payment Confirmation</h1><p>Your payment for order {{.OrderID}} has status {{.Status}}.</p></body></html>`
	DeliveryAssignmentTemplate = `<!DOCTYPE html>
<html><body><h1>Delivery Assigned</h1><p>Delivery person {{.DeliveryPerson}} has been assigned to order {{.OrderID}}.</p></body></html>`
)
