package usecase

import (
	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/usecase/dto"
	"github.com/lee212400/myProject/usecase/repository"
)

type OrderInteractor struct {
	orderRepository repository.OrderRepository
}

func NewOrderInteractor(orderRepo repository.OrderRepository) *OrderInteractor {
	return &OrderInteractor{orderRepository: orderRepo}
}

func (o *OrderInteractor) Save(in *dto.PostInputDto) error {
	order, err := entity.NewOrder(in.Order.Id, in.Order.UserId, in.Order.Items)
	if err != nil {
		return err
	}

	err = o.orderRepository.Save(order)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderInteractor) FindByID(id string) (*entity.Order, error) {
	foundOrder, err := o.orderRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundOrder, nil
}
