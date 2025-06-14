package service

import (
	"context"
	"go-learn-rest-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryid int)
	FindById(ctx context.Context, categoryid int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
