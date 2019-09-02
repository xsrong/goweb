package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres password=jxaf00504771 dbname=weibo_app sslmode=disable")
	if err != nil {
		panic(err)
	}
	DB.LogMode(true)
	DB.AutoMigrate(&User{}, &Relationship{})
}
