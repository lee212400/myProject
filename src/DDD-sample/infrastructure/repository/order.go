package repository

import (
	"myProject/src/DDD-sample/domain/model"
)

type OrderRepository struct {
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (o *OrderRepository) Save(order *model.Order) error {
	// db処理
	return nil
}

func (o *OrderRepository) FindByID(id string) (*model.Order, error) {
	order := &model.Order{}

	// db処理

	return order, nil
}
