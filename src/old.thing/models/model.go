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
	Name     string // column name will be `name`
	Username string `json:"username" form:"username"`
	Email    string // column name will be `email`
	Password string `json:"password" form:"password"`
	Birthday time.Time
	Avatar   string
	Address  string // column name will be `Address`
	Lat      float64
	Lng      float64
}

// func Table name
func (User) TableName() string {
	return "users"
}

// Post struct
type Post struct {
	Model
	Title       string
	UserId      int
	User        User
	Category    Category
	CategoryId  int
	Description string
	Contact     string
	Lat         float64
	Lng         float64
}

func (Post) TableName() string {
	return "posts"
}

// Category struct
type Category struct {
	Model
	Name        string
	Description string // column name will be `birthday`
	Image       string // column name will be `Avatar`
}

func (Category) TableName() string {
	return "categories"
}
