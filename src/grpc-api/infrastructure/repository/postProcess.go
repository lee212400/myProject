package repository

import (
	"fmt"

	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/infrastructure/db"
)

type PostProcessRepositoryImpl struct{}

func NewPostProcessRepositoryImpl() *PostProcessRepositoryImpl {
	return &PostProcessRepositoryImpl{}
}

func (i *PostProcessRepositoryImpl) PostProcess(ctx *entity.Context, err *error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	db.CloseDb(ctx, *err == nil)
}
