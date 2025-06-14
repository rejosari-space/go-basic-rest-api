package controller

import (
	"go-learn-rest-api/service"
	"net/http"
)

type CategoryControllerImpl struct {
	categoryService service.CategoryService
}
func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params http.RoundTripper.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

