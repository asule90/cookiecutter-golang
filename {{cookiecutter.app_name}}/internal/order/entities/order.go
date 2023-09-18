package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	ID            string
	Name          string
	Channel       string
	IsPickup      bool
	DateShipment  time.Time
	ShipTime      string
	Status        string
	PaymentMethod string
	PaymentStatus string
	Note          string
	Promo         string
	GrandTotal    decimal.Decimal
	CreatedAt     time.Time
	CreatedBy     string
	UpdatedAt     *time.Time
	UpdatedBy     string
}

type OrderItem struct {
	ID      string
	OrderID string
	SKUID   string
	SKUCode string
	Qty     int
	Price   decimal.Decimal
	Promo   string
	Note    string
}
