package entity

import "context"

type Context struct {
	context.Context
	Response any
	Session  map[string]any
}
