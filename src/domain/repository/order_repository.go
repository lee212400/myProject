package repository

import (
	"myProject/src/domain/model"
)

// ビジネスルールで決まっている機能

type OrderRepository interface {
	Save(order *model.Order) error
	FindByID(id string) (*model.Order, error)
}
