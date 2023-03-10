package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Product struct {
	ID          int    `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:text"`
	Price       int    `gorm:"type:integer"`
	Image       string `gorm:"type:text"`
}

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123@tcp(localhost:3306)/products?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Product{})
}