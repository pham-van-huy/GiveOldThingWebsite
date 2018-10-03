package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (p *GeoPoint) String() string {
	return fmt.Sprintf("POINT(%v, %v", p.Lat, p.Lng)
}

type User struct {
	ID        uint      // column name will be `id`
	Name      string    // column name will be `name`
	Email     string    // column name will be `email`
	Password  string    // column name will be `Password`
	Birthday  time.Time // column name will be `birthday`
	Avatar    string    // column name will be `Avatar`
	Address   string    // column name will be `Address`
	Location  GeoPoint  `json:"location"`
	CreatedAt time.Time // column name will be `created_at`
}

// set User's table name to be `profiles`
func (User) TableName() string {
	return "users"
}

func main() {
	db, err := gorm.Open("mysql", "root:@/oldthing?charset=utf8&parseTime=True&loc=Local")

	fmt.Println("err", err)
	defer db.Close()

	db.AutoMigrate(&User{})

	// user := User{Name: "Ngoc", Birthday: time.Now()}
	// db.Create(&user)

}
