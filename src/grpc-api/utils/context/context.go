package context

import (
	"context"

	"github.com/google/uuid"
	"github.com/lee212400/myProject/domain/entity"
)

type contextKey string

const (
	requestId contextKey = "request-id"
	traceId   contextKey = "trace-id"
)

func NewContext(ctx context.Context) *entity.Context {
	id := uuid.New()
	ctx = context.WithValue(ctx, requestId, id.String())
	ctx = context.WithValue(ctx, traceId, id.String())

	return &entity.Context{
		Context: ctx,
		Session: map[string]any{},
	}
}
