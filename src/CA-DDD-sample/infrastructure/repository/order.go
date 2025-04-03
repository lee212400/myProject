package repository

import (
	"github.com/lee212400/myProject/domain/entity"
)

type OrderRepository struct {
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (o *OrderRepository) Save(order *entity.Order) error {
	// db処理
	return nil
}

func (o *OrderRepository) FindByID(id string) (*entity.Order, error) {
	order := &entity.Order{}

	// db処理

	return order, nil
}
