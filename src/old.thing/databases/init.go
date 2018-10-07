package databases

import (
	"fmt"

	m "old.thing/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() {
	db, err := gorm.Open("mysql", "root:@/oldthing?charset=utf8&parseTime=True&loc=Local")

	fmt.Println("err", err)
	defer db.Close()

	db.AutoMigrate(&m.User{})
	db.AutoMigrate(&m.Post{})
	db.AutoMigrate(&m.Category{})
	db.AutoMigrate(&m.Image{})
	// user := User{Name: "Ngoc", Birthday: time.Now()}
	// db.Create(&user)
}
