package helper

import (
	"go-learn-rest-api/model/domain"
	"go-learn-rest-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryresponses(categories []domain.Category) []web.CategoryResponse {
	var categoryresponses []web.CategoryResponse
	for _, category := range categories {
		categoryresponses = append(categoryresponses, ToCategoryResponse(category))
	}

	return categoryresponses
}
