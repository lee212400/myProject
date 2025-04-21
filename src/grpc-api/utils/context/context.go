package context

import (
	"context"

	"github.com/google/uuid"
	"github.com/lee212400/myProject/domain/entity"
)

func NewContext(ctx context.Context) *entity.Context {
	id := uuid.New()
	ctx = context.WithValue(ctx, "request-id", id.String())
	ctx = context.WithValue(ctx, "trace-id", id.String())

	return &entity.Context{
		Context: ctx,
		Session: map[string]any{},
	}
}
