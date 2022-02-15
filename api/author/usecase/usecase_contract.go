package usecase

import (
	"go-library/api/author/repository"
	"gorm.io/gorm"
)

type UsecaseContract interface {
	AuthorContract
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
