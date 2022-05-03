package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Global vars
var (
	db *gorm.DB
)

/**
 * Establish Connection to DB
 */
func Connect() {
	// DB connection
	// Format: username:password/db_name
	d, err := gorm.Open("mysql", "pnguyen:deptraithimoiconhieuduaiu@tcp(<hostname>)/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
