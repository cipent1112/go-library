package entity

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Author struct {
	ID   string `json:"id" gorm:"column:id; PRIMARY_KEY"`
	Name string `json:"name" validate:"required" gorm:"column:name;"`
}

func (Author) TableName() string {
	return "author"
}

func (b *Author) BeforeCreate(*gorm.DB) (err error) {
	b.ID = uuid.NewV4().String()
	return
}
