package repository

import (
	"github.com/lee212400/myProject/domain/entity"
)

type OrderRepository interface {
	Save(order *entity.Order) error
	FindByID(id string) (*entity.Order, error)
}
