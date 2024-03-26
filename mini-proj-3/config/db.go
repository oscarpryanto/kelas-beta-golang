package config

import (
	"fmt"
	"log"
	"miniproject3/manajemen_buku_perpustakaan/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlDb struct {
	DB *gorm.DB
}

var Mysql MySqlDb

func OpenDb() {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	mysqlConn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Mysql = MySqlDb{
		DB: mysqlConn,
	}

	err = autoMigrate(mysqlConn)
	if err != nil {
		log.Fatal(err)
	}
}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Book{},
	)

	if err != nil {
		log.Fatal(err)
	}

	return nil

}
