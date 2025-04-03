package usecase

import (
	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/usecase/dto"
)

type OrderService interface {
	Save(in *dto.PostInputDto) error
	FindByID(id string) (*entity.Order, error)
}
