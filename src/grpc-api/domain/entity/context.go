package entity

import "context"

type Context struct {
	ctx      context.Context
	Response any
	Session  map[string]any
}
