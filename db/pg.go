package db

import (
	"log"

	"github.com/hecterbonha/sergeant/config"
	"github.com/hecterbonha/sergeant/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB

func ConnectionInit() {
	dbURL := config.PostgresURL()

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.User{}, &model.Profile{}, &model.Credentials{})
	Postgres = db
}
