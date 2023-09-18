package models

import (
	"database/sql"
	"time"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order/entities"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/null"
	"github.com/jinzhu/copier"
	"github.com/shopspring/decimal"
)

type Orders struct {
	ID            uuidv7
	Name          string
	Channel       ChannelEnum
	IsPickup      bool
	DateShipment  time.Time
	ShipTime      *ShipmentTime
	Status        *OrderStatus
	PaymentMethod PaymentMethodEnum
	PaymentStatus PaymentStatusEnum
	Note          sql.NullString
	Promo         sql.NullString
	GrandTotal    decimal.Decimal `gorm:"type:decimal(10,2)"`
	CreatedAt     time.Time
	CreatedBy     string
	UpdatedAt     sql.NullTime
	UpdatedBy     string
}

type OrderItems struct {
	ID      uuidv7
	OrderID uuidv7
	SKUID   uuidv7
	SKUCode string
	Qty     int
	Price   decimal.Decimal `gorm:"type:decimal(10,2)"`
	Promo   sql.NullString
	Note    sql.NullString
}

type OrderStatus string

const (
	Pending    OrderStatus = "pending"
	Diproduksi OrderStatus = "diproduksi"
	Ready      OrderStatus = "ready"
	Dikirim    OrderStatus = "dikirim"
	Diterima   OrderStatus = "diterima"
)

type ShipmentTime string

const (
	Pagi  ShipmentTime = "pagi"
	Siang ShipmentTime = "siang"
	Sore  ShipmentTime = "sore"
)

type PaymentStatusEnum string

const (
	Belum     PaymentStatusEnum = "belum"
	Lunas     PaymentStatusEnum = "lunas"
	TrialFree PaymentStatusEnum = "trial/free"
)

type PaymentMethodEnum string

const (
	CashToko  PaymentMethodEnum = "Cash Toko"
	CashTitip PaymentMethodEnum = "Cash Titip"
	Transfer  PaymentMethodEnum = "Transfer"
	Qris      PaymentMethodEnum = "Qris"
	Debit     PaymentMethodEnum = "Debit"
)

type ChannelEnum string

const (
	Toko ChannelEnum = "toko"
	IGWA ChannelEnum = "IG/WA"
)

type ConsigneeEnum string

const (
	B2B   ConsigneeEnum = "B2B"
	B2C   ConsigneeEnum = "B2C"
	Motor ConsigneeEnum = "motor"
)

type uuidv7 string

func (o Orders) ToEntity() entities.Order {
	entAppUser := entities.Order{}
	_ = copier.Copy(&entAppUser, &o)

	entAppUser.CreatedAt = o.CreatedAt
	entAppUser.UpdatedAt = null.NullTimeToTimePtr(o.UpdatedAt)

	return entAppUser
}
