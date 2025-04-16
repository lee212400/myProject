package presenter

import "github.com/lee212400/myProject/usecase"

type UserPresenter struct {
	outputPort *usecase.UserOutputPort
}

func NewUsePresenter() *UserPresenter {
	return &UserPresenter{}
}
