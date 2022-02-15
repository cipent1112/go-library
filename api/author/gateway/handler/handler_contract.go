package handler

import "go-library/api/author/usecase"

type HandlerContract interface {
	AuthorContract
}

type Handler struct {
	Usecase usecase.UsecaseContract
	Name    string
}

func NewHandler(uc usecase.UsecaseContract) HandlerContract {
	return &Handler{
		Usecase: uc,
		Name:    "Author",
	}
}
