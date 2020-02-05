package commons

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gokit-poc/models"
)

var (
	GlobalDB *gorm.DB
	err      error
)

func CreateDatabase(dialect, uri string) *gorm.DB {
	println("Creating DB")
	GlobalDB, err = gorm.Open(dialect, uri)
	if err != nil {
		panic("Error when creating database: " + err.Error())
	}

	println("Migrating tables")
	migrate(GlobalDB)

	return GlobalDB
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Account{})
}
