package config

import (
	"dot/app/models/domain"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var e error

func Connection() (*gorm.DB, error) {
	user := "root"
	password := "Buat$and1"
	host := "127.0.0.1"
	port := "3306"
	dbname := "dot_rest_api"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	DB, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
	e = DB.AutoMigrate(&domain.User{}, &domain.Book{}, &domain.Category{})
	if e != nil {
		return nil, e
	}
	return DB, nil

}
