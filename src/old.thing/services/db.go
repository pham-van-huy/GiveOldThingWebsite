package services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func DB_Instance() *gorm.DB {
	if db == nil {
		db = DB_Open()
	}

	return db
}

func DB_Open() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/oldthing?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil
	}

	return db
}

func DB_Close() {
	if db != nil {
		db.Close()
	}
}
