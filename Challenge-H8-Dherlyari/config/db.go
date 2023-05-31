package config

import (
	"Challenge-H8-Dherlyari/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	conn := "host=localhost user=postgres password=root dbname=TrialClassGolangHacktiv8 port=5432"

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Tidak Terkoneksi", err.Error())
	} else {
		fmt.Println("Database Terkoneksi")
		DB = db
	}

	db.AutoMigrate(&entity.Users{})
}
