package repository

import (
	"context"
	"database/sql"
	"go-learn-rest-api/model/domain"
)

type CategoryRepositoryImpl struct {

}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT INTO category(name) values(?)"
	result,err := tx.ExecContext(ctx,query,category.Name)
	if err != nil {
		panic(err)
	}

	id,err:= result.LastInsertId()
	if err != nil {
		panic(err)
	}
	category.Id=int(id)
	return category

}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, ctaegory domain.Category) domain.Category {
	panic("not implemented") // TODO: Implement
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

