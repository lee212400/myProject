package repository

import "github.com/lee212400/myProject/domain/entity"

type PostProcessRepository interface {
	PostProcess(ctx *entity.Context, err *error)
}
