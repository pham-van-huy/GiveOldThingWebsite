package models

import (
	"fmt"
	"time"
)

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (p *GeoPoint) String() string {
	return fmt.Sprintf("POINT(%v, %v)", p.Lat, p.Lng)
}

// User model
type User struct {
	ID        string    `json:"uuid" form:"-"`
	Name      string    // column name will be `name`
	Username  string    `json:"username" form:"username"`
	Email     string    // column name will be `email`
	Password  string    `json:"password" form:"password"`
	Birthday  time.Time // column name will be `birthday`
	Avatar    string    // column name will be `Avatar`
	Address   string    // column name will be `Address`
	Location  GeoPoint  `json:"location"`
	CreatedAt time.Time // column name will be `created_at`
}

// func Table name
func (User) TableName() string {
	return "users"
}

// Post struct
type Post struct {
	ID          uint // column name will be `id`
	Title       string
	UserId      int
	Owner       User
	Category    Category
	CategoryId  int       // column name will be `Password`
	Description string    // column name will be `birthday`
	Contact     string    // column name will be `Avatar`
	Location    GeoPoint  `json:"location"`
	CreatedAt   time.Time // column name will be `created_at`
	UpdatedAt   time.Time // column name will be `created_at`
}

func (Post) TableName() string {
	return "posts"
}

// Category struct
type Category struct {
	ID          uint // column name will be `id`
	Name        string
	Description string    // column name will be `birthday`
	Image       string    // column name will be `Avatar`
	CreatedAt   time.Time // column name will be `created_at`
	UpdatedAt   time.Time // column name will be `created_at`
}

func (Category) TableName() string {
	return "categories"
}
