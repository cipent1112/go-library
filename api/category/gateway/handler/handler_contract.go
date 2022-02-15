package handler

import "go-library/api/category/usecase"

type HandlerContract interface {
	CategoryContract
}

type Handler struct {
	Usecase usecase.UsecaseContract
	Name    string
}

func NewHandler(uc usecase.UsecaseContract) HandlerContract {
	return &Handler{
		Usecase: uc,
		Name:    "Category",
	}
}
