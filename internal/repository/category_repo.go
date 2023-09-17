package repository

import (
	"context"
	"take-out/common"
	"take-out/internal/api/request"
	"take-out/internal/model"
)

type CategoryRepo interface {
	Insert(ctx context.Context, category model.Category) error
	PageQuery(ctx context.Context, dto request.CategoryPageQueryDTO) (*common.PageResult, error)
}
