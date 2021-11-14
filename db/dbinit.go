package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var Db *gorm.DB
func init()  {
	Db = GormDb()
}

func GormDb() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/testgo?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil {
		log.Fatal(err)
	}

	sqlDB,err := db.DB()
	if err !=nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}