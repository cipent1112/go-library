package migrate

import (
	"gorm.io/gorm"
)

/*	Do Check and migrate if table is not exist */
func AutoMigrate(db *gorm.DB) {
	MigrateAuthor(db)
	MigrateCategory(db)
	MigrateBook(db)
}
