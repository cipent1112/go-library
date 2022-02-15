package usecase

import (
	"go-library/api/category/repository"
	"gorm.io/gorm"
)

type UsecaseContract interface {
	CategoryContract
}

type Usecase struct {
	DB   *gorm.DB
	Repo repository.RepositoryContract
}

func NewUsecase(db *gorm.DB, repo repository.RepositoryContract) UsecaseContract {
	return &Usecase{
		DB:   db,
		Repo: repo,
	}
}
