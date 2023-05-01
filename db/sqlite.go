package db

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDatabase struct {
	DB *gorm.DB
}

func (s *SqliteDatabase) GetDB() *gorm.DB {
	return s.DB
}

type SqliteConfig struct {
	Path string `default:"db/development.db"`
}

func NewSqliteConfig() *SqliteConfig {
	var config SqliteConfig
	envconfig.Process("db", &config)
	return &config
}

func (sc *SqliteConfig) ToString() string {
	return fmt.Sprintf("%s", sc.Path)
}

func NewSqlite(models []interface{}) *SqliteDatabase {
	config := NewSqliteConfig()

	if config.Path == "db/development.db" {
		os.MkdirAll("db", os.ModePerm)
	}

	db, err := gorm.Open(sqlite.Open(config.ToString()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &SqliteDatabase{DB: db}
}
