package repository

import (
	"github.com/lee212400/myProject/domain/model"
)

// ビジネスルールで決まっている機能

type OrderRepository interface {
	Save(order *model.Order) error
	FindByID(id string) (*model.Order, error)
}
