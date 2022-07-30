package migrations

import (
	"github.com/marcellorhcp/webapi-mvc/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
