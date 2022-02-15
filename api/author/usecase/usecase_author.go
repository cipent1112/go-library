package usecase

import (
	"go-library/api/author/entity"
	"go-library/lib/presenter"
	"go-library/lib/request"
	"math"
)

type AuthorContract interface {
	Get(id string) (res entity.Author, err error)
	GetList(param request.List) (res []entity.Author, meta presenter.Meta, err error)
	Store(input entity.Author) (err error)
}

func (u Usecase) Get(id string) (res entity.Author, err error) {
	res, err = u.Repo.Get(u.DB, id)
	if err != nil {
		return
	}

	return
}

func (u Usecase) GetList(param request.List) (res []entity.Author, meta presenter.Meta, err error) {
	res, total, err := u.Repo.GetList(u.DB, param)
	if err != nil {
		return
	}

	meta = presenter.Meta{
		TotalData: int(total),
		TotalPage: int(math.Ceil(float64(total) / float64(param.PerPage))),
		Page:      param.Page,
		PerPage:   param.PerPage,
	}

	return
}

func (u Usecase) Store(input entity.Author) (err error) {
	db := u.DB.Begin()

	if err = u.Repo.Store(db, input); err != nil {
		db.Rollback()
		return
	}

	return db.Commit().Error
}
