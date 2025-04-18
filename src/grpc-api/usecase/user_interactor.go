package usecase

import (
	"log"

	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/usecase/dto"
	"github.com/lee212400/myProject/usecase/repository"
)

type UserInteractor struct {
	postProcessRepository repository.PostProcessRepository
	userRepository        repository.UserRepository
	outputPort            UserOutputPort
}

func NewUserInteractor(postProcessRepository repository.PostProcessRepository, userRepository repository.UserRepository, outputPort UserOutputPort) *UserInteractor {
	return &UserInteractor{
		postProcessRepository: postProcessRepository,
		userRepository:        userRepository,
		outputPort:            outputPort,
	}
}

func (i *UserInteractor) GetUser(ctx *entity.Context, in *dto.GetUserInputDto) (err error) {
	log.Println("Interactor GetUser")
	defer i.postProcessRepository.PostProcess(ctx, &err)
	u, err := i.userRepository.GetUser(ctx, in.UserId)
	if err != nil {
		return
	}

	err = i.outputPort.GetUser(ctx, &dto.GetUserOutputDto{
		User: u,
	})

	return
}
func (i *UserInteractor) CreateUser(ctx *entity.Context, in *dto.CreateUserInputDto) (err error) {
	log.Println("Interactor CreateUser")
	defer i.postProcessRepository.PostProcess(ctx, &err)
	uId, err := i.userRepository.CreateUser(ctx, in.User.FirstName, in.User.LastName, in.User.Email, in.User.Age)
	if err != nil {
		return
	}

	u, err := i.userRepository.GetUser(ctx, uId)
	if err != nil {
		return
	}

	err = i.outputPort.CreateUser(ctx, &dto.CreateUserOutputDto{
		User: u,
	})

	return
}
func (i *UserInteractor) UpdateUser(ctx *entity.Context, in *dto.UpdateUserInputDto) (err error) {
	defer i.postProcessRepository.PostProcess(ctx, &err)

	err = i.userRepository.UpdateUser(ctx, in.UserId, in.Age)
	if err != nil {
		return
	}

	u, err := i.userRepository.GetUser(ctx, in.UserId)
	if err != nil {
		return
	}

	err = i.outputPort.UpdateUser(ctx, &dto.UpdateUserOutputDto{
		User: u,
	})

	return
}
func (i *UserInteractor) DeleteUser(ctx *entity.Context, in *dto.DeleteUserInputDto) (err error) {
	defer i.postProcessRepository.PostProcess(ctx, &err)

	err = i.userRepository.DeleteUser(ctx, in.UserId)
	if err != nil {
		return
	}

	err = i.outputPort.DeleteUser(ctx, &dto.DeleteUserOutputDto{})

	return
}
