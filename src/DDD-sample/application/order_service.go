package application

import (
	"github.com/lee212400/myProject/domain/model"
	"github.com/lee212400/myProject/domain/repository"
)

type OrderRepository interface {
	Save(order *model.Order) error
	FindByID(id string) (*model.Order, error)
}

type OrderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) *OrderService {
	return &OrderService{orderRepository: orderRepo}
}

func (o *OrderService) Save(orderId string, userId string, items []*model.OrderItem) error {
	order, err := model.NewOrder(orderId, userId, items)
	if err != nil {
		return err
	}

	err = o.orderRepository.Save(order)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderService) FindByID(id string) (*model.Order, error) {
	foundOrder, err := o.orderRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundOrder, nil
}
