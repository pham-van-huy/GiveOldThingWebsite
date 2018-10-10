package databases

import (
	"fmt"

	"github.com/icrowley/fake"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	m "old.thing/models"
)

func Init() {
	db, err := gorm.Open("mysql", "root:@/oldthing?charset=utf8&parseTime=True&loc=Local")

	fmt.Println("err", err)
	defer db.Close()

	// redis := s.Redis_Connect()

	db.AutoMigrate(&m.User{})
	user1 := m.User{Name: "Phan Ngoc", Username: "phann123", Password: "123", Address: "Da Nang"}
	db.Create(&user1)

	db.AutoMigrate(&m.Post{})

	db.AutoMigrate(&m.Category{})
	cate1 := m.Category{Name: "Xe co"}
	cate2 := m.Category{Name: "Do dien tu"}
	cate3 := m.Category{Name: "Bat dong san"}
	db.Create(&cate1)
	db.Create(&cate2)
	db.Create(&cate3)

	db.AutoMigrate(&m.Image{})

	db.AutoMigrate(&m.City{})
	city1 := m.City{Name: "Da Nang"}
	city2 := m.City{Name: "Quang Nam"}
	city3 := m.City{Name: "Ha Noi"}
	city4 := m.City{Name: "Ho Chi Minh"}
	city5 := m.City{Name: fake.City()}
	db.Create(&city1)
	db.Create(&city2)
	db.Create(&city3)
	db.Create(&city4)
	db.Create(&city5)
}
