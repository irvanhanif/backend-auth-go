package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB() {
	var dsn string
	var err error
	switch os.Getenv("DB_TYPE") {
	case "mysql":
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"), 
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	case "pg":
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", 
			os.Getenv("DB_HOST"), 
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "auth.",
				SingularTable: true,
			},
		})

	default:
		panic("Failed to connect DB")
	}

	if err != nil {
		panic("Failed to connect DB")
	}

}