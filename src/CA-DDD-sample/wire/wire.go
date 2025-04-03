//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/lee212400/myProject/infrastructure/repository"
	"github.com/lee212400/myProject/interface/handler"
	"github.com/lee212400/myProject/usecase"
	repo "github.com/lee212400/myProject/usecase/repository"
)

func InitOrder() *handler.OrderHandler {
	wire.Build(
		usecase.NewOrderInteractor,
		repository.NewOrderRepositoryImpl,
		handler.NewOrderHandler,
		// repository DI(interface -> 実装)
		wire.Bind(new(repo.OrderRepository), new(*repository.OrderRepositoryImpl)),
		// usecase DI(interface -> 実装)
		wire.Bind(new(usecase.OrderService), new(*usecase.OrderInteractor)),
	)
	return nil
}
