package repository

import (
	"github.com/lee212400/myProject/domain/entity"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepositoryImpl() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{}
}

func (o *OrderRepositoryImpl) Save(order *entity.Order) error {
	// db処理
	return nil
}

func (o *OrderRepositoryImpl) FindByID(id string) (*entity.Order, error) {
	order := &entity.Order{}

	// db処理

	return order, nil
}
