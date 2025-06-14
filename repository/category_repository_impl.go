package repository

import (
	"context"
	"database/sql"
	"go-learn-rest-api/helper"
	"go-learn-rest-api/model/domain"
)

type CategoryRepositoryImpl struct {

}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT INTO category(name) values(?)"
	result,err := tx.ExecContext(ctx,query,category.Name)
	helper.PanicIfError(err)

	id,err:= result.LastInsertId()
	helper.PanicIfError(err)


	category.Id=int(id)
	return category

}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query:= "UPDATE category SET name =? WHERE id = ?"
	_,err:= tx.ExecContext(ctx,query ,category.Name,category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	panic("not implemented") // TODO: Implement
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) domain.Category {
	panic("not implemented") // TODO: Implement
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	panic("not implemented") // TODO: Implement
}

