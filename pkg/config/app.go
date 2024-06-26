package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB

func Connect() {
	err := godotenv.Load()
	POSTGRESDB := os.Getenv("POSTGRESDB")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	d, err := gorm.Open("postgres", POSTGRESDB)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
