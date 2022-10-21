package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabaseConnection(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		fmt.Errorf("Error when connect to database : %v", err.Error())
		return nil, err
	}

	return db, nil
}
