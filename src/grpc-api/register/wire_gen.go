// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package register

import (
	"github.com/lee212400/myProject/infrastructure/db"
	"github.com/lee212400/myProject/infrastructure/repository"
	"github.com/lee212400/myProject/interface/controller"
	"github.com/lee212400/myProject/interface/presenter"
	"github.com/lee212400/myProject/usecase"
)

// Injectors from wire.go:

func UserInit() *controller.UserController {
	postProcessRepositoryImpl := repository.NewPostProcessRepositoryImpl()
	sqlDB := db.NewDb()
	userRepositoryImpl := repository.NewUserRepositoryImpl(sqlDB)
	userPresenter := presenter.NewUserPresenter()
	userInteractor := usecase.NewUserInteractor(postProcessRepositoryImpl, userRepositoryImpl, userPresenter)
	userController := controller.NewUserController(userInteractor)
	return userController
}
