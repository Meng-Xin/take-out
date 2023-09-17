package repository

import (
	"context"
	"take-out/internal/model"
)

type CategoryRepo interface {
	Insert(ctx context.Context, category model.Category) error
}
