package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	url, ok := viper.Get("DATABASE_URL").(string)
	if !ok {
		log.Panicln(fmt.Errorf("could not find database url"))
	}
	connection, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	db = connection
	return db, err
}

func GetDB() *gorm.DB {
	return db
}
