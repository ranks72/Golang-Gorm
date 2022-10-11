package database

import (
	"fmt"
	"log"

	//"tugas_2/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "" //silahkan masukan user db anda
	password = "" //silahkan masukan pass db anda
	dBport   = "5432"
	dBname   = "belajar_gorm"
	db       *gorm.DB
	err      error
)

func StartDb() {
	//fmt.Printf("postgresql:user:%s, password:%s dB:@%s:%s/%s sslmode=disable", user, password, host, dBport, dBname)
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dBname, dBport)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err.Error())
	}

	if err != nil {
		log.Fatal("error while tyring to ping the database connection", err.Error())
	}

	fmt.Println("successfully connected to my database")

	//untuk buat tabel baru
	//db.Debug().AutoMigrate(models.Order{}, models.Item{})

}

func GetDb() *gorm.DB {
	return db
}
