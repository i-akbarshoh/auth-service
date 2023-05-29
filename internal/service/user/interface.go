package user

import (
	"context"
)

type Repository interface {
	Create(context.Context, Create) error
	Get(context.Context, string) (Get, error)
}
