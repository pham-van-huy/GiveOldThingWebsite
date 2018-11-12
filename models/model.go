package models

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (p *GeoPoint) String() string {
	return fmt.Sprintf("POINT(%v, %v)", p.Lat, p.Lng)
}

// Scan implements the Scanner interface which will scan the postgis POINT(x, y) into the GeoPoint struct
func (p *GeoPoint) Scan(val interface{}) error {
	b, err := hex.DecodeString(string(val.([]uint8)))
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	var wkbByteOrder uint8
	if err := binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
		return err
	}

	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case 0:
		byteOrder = binary.BigEndian
	case 1:
		byteOrder = binary.LittleEndian
	default:
		return fmt.Errorf("Invalid byte order %u", wkbByteOrder)
	}

	var wkbGeometryType uint64
	if err := binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
		return err
	}

	if err := binary.Read(r, byteOrder, p); err != nil {
		return err
	}

	return nil
}

func (p *GeoPoint) Value() (driver.Value, error) {
	return p.String(), nil
}

type Model struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// User model
type User struct {
	Model
	Name     string    `json:"Name" form:"Name"`
	Username string    `json:"Username" form:"Username"`
	Email    string    `json:"Email" form:"Email"`
	Password string    `json:"Password" form:"Password"`
	Birthday time.Time `json:"Birthday" form:"Birthday"`
	Avatar   string    `json:"Avatar" form:"Avatar"`
	Address  string    `json:"Address" form:"Address"`
	Lat      float64   `json:"Lat" form:"Lat"`
	Lng      float64   `json:"Lng" form:"Lng"`
}

// func Table name
func (User) TableName() string {
	return "users"
}

// Post struct
type Post struct {
	Model
	Title       string `json:"Title" form:"Title" validate:"min=3,max=100"`
	UserId      int    `json:"UserId" form:"UserId" validate:nonzero`
	User        User
	Category    Category
	CategoryId  int `json:"CategoryId" form:"CategoryId" validate:nonzero`
	CityId      int `json:"CityId" form:"CityId" validate:nonzero`
	City        City
	Price       int     `json:"Price" form:"Price"`
	Images      []Image `gorm:"foreignkey:OwnerId"`
	Description string  `json:"Description" form:"Description" validate:"min=3,max=1000"`
	Contact     string  `json:"Contact" form:"Contact"`
	Lat         float64 `json:"Lat" form:"Lat"`
	Lng         float64 `json:"Lng" form:"Lng"`
	Action      int     `json:"Action" form:"Action"`
}

func (Post) TableName() string {
	return "posts"
}

// Category struct
type Category struct {
	Model
	Name        string `json:"Name" form:"Name"`
	Description string `json:"Description" form:"Description"`
	Image       Image  `gorm:"polymorphic:Owner;"`
}

func (Category) TableName() string {
	return "categories"
}

// Post struct
type Image struct {
	Model
	OwnerId     int
	OwnerType   string
	Name        string `json:"Name" form:"Name"`
	Description string `json:"Description" form:"Description"`
	Link        string `json:"Link" form:"Link"`
}

func (Image) TableName() string {
	return "images"
}

// City struct
type City struct {
	Model
	Name string `json:"Name" form:"Name"`
}

func (City) TableName() string {
	return "cities"
}

// Post struct
type BookMark struct {
	Model
	UserId int `json:"UserId" form:"UserId" mapstructure:"UserId"`
	PostId int `json:"PostId" form:"PostId" mapstructure:"PostId"`
}

func (BookMark) TableName() string {
	return "bookmarks"
}

const (
	TYPE_POST string = "posts"
)
