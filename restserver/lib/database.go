package lib

import (
	"github.com/muhaobing-eng/std-go/go-common/database"
	"github.com/muhaobing-eng/std-go/restserver/config"

	"gorm.io/gorm"
)

var (
	dbSingleInstance *gorm.DB
)

func InitDatabase(cfg config.DatabaseConfig) error {
	db, err := database.New(cfg.GetDatabaseOption())
	if err != nil {
		return err
	}
	dbSingleInstance = db
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return dbSingleInstance
}
