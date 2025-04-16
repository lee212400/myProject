package controller

import "github.com/lee212400/myProject/usecase"

type UserController struct {
	inputPort *usecase.UserInputPort
}

func NewUserController() *UserController {
	return &UserController{}
}
