package context

import (
	"context"

	"github.com/lee212400/myProject/domain/entity"
)

func NewContext(ctx context.Context) *entity.Context {
	return &entity.Context{
		Ctx:     ctx,
		Session: map[string]any{},
	}
}
