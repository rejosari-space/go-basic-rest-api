package service

import (
	"database/sql"
	"go-learn-rest-api/repository"
)

type CategoryServiceImpl struct {
	db         *sql.DB
	repository repository.CategoryRepository
}
