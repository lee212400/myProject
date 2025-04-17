package usecase

import (
	"fmt"

	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/usecase/dto"
	"github.com/lee212400/myProject/usecase/repository"
)

type UserInteractor struct {
	userRepository repository.UserRepository
	outputPort     UserOutputPort
}

func NewUserInteractor(userRepository repository.UserRepository, outputPort UserOutputPort) *UserInteractor {
	return &UserInteractor{
		userRepository: userRepository,
		outputPort:     outputPort,
	}
}

func (i *UserInteractor) GetUser(ctx *entity.Context, in *dto.GetUserInputDto) (err error) {
	fmt.Println("Interactor GetUser")
	u, _ := i.userRepository.GetUser(ctx, in.UserId)
	_ = i.outputPort.GetUser(ctx, &dto.GetUserOutputDto{
		User: u,
	})

	return
}
func (i *UserInteractor) CreateUser(ctx *entity.Context, in *dto.CreateUserInputDto) (err error) {
	return
}
func (i *UserInteractor) UpdateUser(ctx *entity.Context, in *dto.UpdateUserInputDto) (err error) {
	return
}
func (i *UserInteractor) DeleteUser(ctx *entity.Context, in *dto.DeleteUserInputDto) (err error) {
	return
}
