package migrations

import (
	"github.com/LenilsonTeixeira/books-api/database/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
