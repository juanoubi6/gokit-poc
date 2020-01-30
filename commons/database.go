package commons

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gokit-poc/models"
	"os"
	"path"
)

func CreateDatabase(uri string) *gorm.DB {
	println("Creating DB")
	wd, _ := os.Getwd()
	db, err := gorm.Open("sqlite3", path.Join(wd, uri))
	if err != nil {
		panic("Error when creating database: " + err.Error())
	}

	println("Migrating tables")
	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
