package entity

import "context"

type Context struct {
	Ctx      context.Context
	Response any
	Session  map[string]any
}
