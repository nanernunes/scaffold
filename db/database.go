package db

import (
	"log"

	"gorm.io/gorm"
)

type Driver int

const (
	Sqlite Driver = iota
	Postgres
)

type Database interface {
	GetDB() *gorm.DB
}

func New(driver Driver, models []interface{}) (database Database) {
	switch driver {
	case Sqlite:
		database = NewSqlite(models)

	case Postgres:
		database = NewPostgres(models)

	default:
		log.Fatal("not implemented driver")
	}

	database.
		GetDB().
		Callback().
		Create().
		Before("gorm:create").
		Register("update_uuid", GenerateUUID)

	for _, model := range models {
		database.GetDB().AutoMigrate(model)
	}

	return
}
