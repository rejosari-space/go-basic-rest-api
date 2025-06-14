package service

import (
	"context"
	"database/sql"
	"go-learn-rest-api/helper"
	"go-learn-rest-api/model/domain"
	"go-learn-rest-api/model/domain/web"
	"go-learn-rest-api/repository"
)

type CategoryServiceImpl struct {
	db         *sql.DB
	repository repository.CategoryRepository
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.repository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.repository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name

	category = service.repository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryid int) {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.repository.FindById(ctx, tx, categoryid)
	helper.PanicIfError(err)

	service.repository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryid int) web.CategoryResponse {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.repository.FindById(ctx, tx, categoryid)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.repository.FindAll(ctx, tx)

	return helper.ToCategoryresponses(categories)
}
