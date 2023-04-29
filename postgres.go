package scaffold

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

type PostgresConfig struct {
	Host     string `default:"localhost"`
	Port     int    `default:"5432"`
	Name     string `default:"development"`
	User     string `default:"postgres"`
	Pass     string `default:""`
	SSLMode  string `default:"disable"`
	TimeZone string `default:"America/Sao_Paulo"`
}

func NewPostgresConfig() *PostgresConfig {
	var config PostgresConfig
	envconfig.Process("db", &config)
	return &config
}

func (r *PostgresConfig) ToString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		r.Host, r.User, r.Pass, r.Name, r.Port, r.SSLMode, r.TimeZone,
	)
}

func NewPostgres(models []interface{}) *Postgres {
	db, err := gorm.Open(postgres.Open(NewPostgresConfig().ToString()), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	for _, model := range models {
		db.AutoMigrate(model)
	}

	return &Postgres{
		DB: db,
	}
}
