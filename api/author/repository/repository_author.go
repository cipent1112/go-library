package repository

import (
	"fmt"
	"go-library/api/author/entity"
	"go-library/lib/conv"
	"go-library/lib/request"
	"gorm.io/gorm"
	"strings"
)

type AuthorContract interface {
	Get(db *gorm.DB, id string) (res entity.Author, err error)
	GetList(db *gorm.DB, param request.List) (res []entity.Author, total int64, err error)
	Store(db *gorm.DB, input entity.Author) (err error)
}

func (repo Repository) Get(db *gorm.DB, id string) (res entity.Author, err error) {
	if err = db.
		Where(entity.Author{ID: id}).
		Where("deleted_at IS NULL").
		Take(&res).Error; err != nil {
		err = fmt.Errorf("author with id %s not found", id)
		return
	}

	return
}

func (repo Repository) GetList(db *gorm.DB, param request.List) (res []entity.Author, total int64, err error) {
	res = []entity.Author{}
	query := db.Model(entity.Author{}).Where("deleted_at IS NULL")

	page := param.PerPage * (param.Page - 1)
	if param.Search != "" {
		formattedTextSearch := "%%" + param.Search + "%%"
		query = query.Where("name LIKE ?", formattedTextSearch)
	}

	if err := query.Count(&total).Error; err != nil {
		return res, total, err
	}

	if total <= 0 {
		return res, total, nil
	}

	var column string
	var order string

	sort := conv.Trim(param.Sort, "-")
	switch sort {
	case "name":
		column = "name"
	default:
		column = "created_at"
	}

	if column != "" {
		isDesc := strings.Contains(param.Sort, "-")
		switch isDesc {
		case true:
			order = "DESC"
		case false:
			order = "ASC"
		}

		query = query.Order(column + " " + order)
	}

	if err := query.
		Limit(param.PerPage).
		Offset(page).
		Find(&res).Error; err != nil {
		return res, total, err
	}

	return
}

func (repo Repository) Store(db *gorm.DB, input entity.Author) (err error) {
	return db.Create(&input).Error
}
