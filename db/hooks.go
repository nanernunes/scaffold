package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GenerateUUID(db *gorm.DB) {
	if db.Statement.Schema != nil {
		for _, field := range db.Statement.Schema.Fields {
			if field.Name == "UUID" {
				field.Set(db.Statement.Context, db.Statement.ReflectValue, uuid.New().String())
			}
		}
	}
}
