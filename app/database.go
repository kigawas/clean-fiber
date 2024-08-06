package app

import (
	"strings"
	"time"

	"github.com/kigawas/clean-fiber/models/domains"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB(dbUrl string, config *gorm.Config) *gorm.DB {
	if strings.HasPrefix(dbUrl, "sqlite://") {
		dbUrl, _ := strings.CutPrefix(dbUrl, "sqlite://")
		return setupDB(dbUrl, sqlite.Open, config)
	} else {
		return setupDB(dbUrl, postgres.Open, config)
	}
}

func setupDB(dbUrl string, open func(string) gorm.Dialector, config *gorm.Config) *gorm.DB {
	_db, err := gorm.Open(open(dbUrl), config)
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := _db.DB()
	if err != nil {
		panic("failed to get connection pool")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute)

	DB = _db
	return _db
}

func GetDB() *gorm.DB {
	return DB
}

func MigrateDB(db *gorm.DB) {
	_ = db.AutoMigrate(&domains.User{})
}
