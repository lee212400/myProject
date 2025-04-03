package entity

import (
	"errors"
	"time"
)

// entity作成・ビジネスロジック作成

type OrderStatus string

const (
	Pending   OrderStatus = "Pending"
	Completed OrderStatus = "Completed"
	Cancelled OrderStatus = "Cancelled"
)

type Order struct {
	Id        string
	UserId    string
	Items     []*OrderItem
	Status    OrderStatus
	CreatedAt time.Time
}

type OrderItem struct {
	ProductID string
	Quantity  int
	Price     float64
}

func NewOrder(id, userId string, items []*OrderItem) (*Order, error) {
	if len(items) == 0 {
		return nil, errors.New("order must have at least one item")
	}
	return &Order{
		Id:        id,
		UserId:    userId,
		Items:     items,
		Status:    Pending,
		CreatedAt: time.Now(),
	}, nil
}

func (o *Order) Complete() {
	o.Status = Completed
}

func (o *Order) Cancel() {
	o.Status = Cancelled
}
