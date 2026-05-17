package infra

import (
	"time"

	"github.com/nats-io/nats.go"
)

func NewNATS(url string) (*nats.Conn, error) {
	return nats.Connect(url,
		nats.Name("food-delivery-broker"),
		nats.MaxReconnects(10),
		nats.ReconnectWait(2*time.Second),
	)
}
