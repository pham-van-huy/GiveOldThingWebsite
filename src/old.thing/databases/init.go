package databases

import (
	"fmt"
	m "models"

	_ "github.com/jilsnzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:@/oldthing?charset=utf8&parseTime=True&loc=Local")

	fmt.Println("err", err)
	defer db.Close()

	db.AutoMigrate(&m.User{})
	db.AutoMigrate(&m.Post{})
	db.AutoMigrate(&m.Category{})
	// user := User{Name: "Ngoc", Birthday: time.Now()}
	// db.Create(&user)

}
