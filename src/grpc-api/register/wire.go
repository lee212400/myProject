//go:build wireinject
// +build wireinject

package register

import (
	"github.com/google/wire"

	"github.com/lee212400/myProject/infrastructure/db"
	"github.com/lee212400/myProject/infrastructure/repository"
	"github.com/lee212400/myProject/interface/controller"
	"github.com/lee212400/myProject/interface/presenter"
	"github.com/lee212400/myProject/usecase"
	uc_repository "github.com/lee212400/myProject/usecase/repository"
)

func UserInit() *controller.UserController {
	wire.Build(
		db.NewDb,
		repository.NewPostProcessRepositoryImpl,
		repository.NewUserRepositoryImpl,
		presenter.NewUserPresenter,
		usecase.NewUserInteractor,
		controller.NewUserController,
		wire.Bind(new(usecase.UserInputPort), new(*usecase.UserInteractor)),
		wire.Bind(new(uc_repository.PostProcessRepository), new(*repository.PostProcessRepositoryImpl)),
		wire.Bind(new(uc_repository.UserRepository), new(*repository.UserRepositoryImpl)),
		wire.Bind(new(usecase.UserOutputPort), new(*presenter.UserPresenter)),
	)
	return nil
}
