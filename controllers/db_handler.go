package controllers

import (
	"database/sql"
	// "fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "os"
)

func connect() *sql.DB {
	// dbHost := os.Getenv("DB_HOST")
	// fmt.Println(dbHost)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func connectGorm() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/db_latihan_pbp2?parseTime=true&loc=Asia%2FJakarta"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}