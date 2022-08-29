package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"wallester_test/entities"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(entities.Customer{})
	log.Println("Database Migration Completed...")
}
